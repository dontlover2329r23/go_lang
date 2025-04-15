package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    n := 10 // по умолчанию
    if len(os.Args) > 1 {
        if val, err := strconv.Atoi(os.Args[1]); err == nil && val >= 10 && val <= 500 {
            n = val
        } else {
            fmt.Println("Invalid matrix size. N must be between 10 and 500.")
            return
        }
    }

    matrix := GenerateMatrix(n)
    if !ValidateMatrix(matrix) {
        fmt.Println("Invalid matrix: must be NxN and non-empty")
        return
    }

    det := DeterminantParallel(matrix)
    fmt.Printf("Determinant: %f\n", det)
}
