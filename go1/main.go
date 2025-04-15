package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите выражение:")
	var input string
	fmt.Scanln(&input)

	result, err := evaluateExpression(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

// evaluateExpression обрабатывает выражение
func evaluateExpression(expr string) (float64, error) {
	tokens := tokenize(expr)
	rpn, err := toRPN(tokens)
	if err != nil {
		return 0, err
	}
	return evalRPN(rpn)
}

// Разбивает выражение на токены
func tokenize(expr string) []string {
	var tokens []string
	var number strings.Builder

	for _, ch := range expr {
		if unicode.IsDigit(ch) || ch == '.' {
			number.WriteRune(ch)
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			if ch == ' ' {
				continue
			}
			tokens = append(tokens, string(ch))
		}
	}
	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}
	return tokens
}

// toRPN: Алгоритм сортировочной станции (Шунтинг-Ярд)
    func toRPN(tokens []string) ([]string, error) {
        var output []string
        var ops []string

        precedence := map[string]int{
            "+": 1,
            "-": 1,
            "*": 2,
            "/": 2,
        }

        for _, token := range tokens {
            switch {
            case isNumber(token):
                output = append(output, token)
            case token == "(":
                ops = append(ops, token)
            case token == ")":
                for len(ops) > 0 && ops[len(ops)-1] != "(" {
                    output = append(output, ops[len(ops)-1])
                    ops = ops[:len(ops)-1]
                }
                if len(ops) == 0 || ops[len(ops)-1] != "(" {
                    return nil, fmt.Errorf("не хватает скобки")
                }
                ops = ops[:len(ops)-1] // Удаляем "("
            default: // оператор
                for len(ops) > 0 && precedence[ops[len(ops)-1]] >= precedence[token] {
                    output = append(output, ops[len(ops)-1])
                    ops = ops[:len(ops)-1]
                }
                ops = append(ops, token)
            }
        }
        for len(ops) > 0 {
            if ops[len(ops)-1] == "(" {
                return nil, fmt.Errorf("не хватает скобки")
            }
            output = append(output, ops[len(ops)-1])
            ops = ops[:len(ops)-1]
        }
        return output, nil
    }

// evalRPN: Считает значение выражения в RPN
func evalRPN(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("недостаточно операндов")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res float64
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("деление на ноль")
				}
				res = a / b
			default:
				return 0, fmt.Errorf("неизвестный оператор: %s", token)
			}
			stack = append(stack, res)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("ошибка вычисления")
	}
	return stack[0], nil
}

// Проверка, является ли строка числом
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
