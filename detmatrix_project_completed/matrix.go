package main

import (
    "math/rand"
    "sync"
)

func GenerateMatrix(n int) [][]float64 {
    m := make([][]float64, n)
    for i := range m {
        m[i] = make([]float64, n)
        for j := range m[i] {
            m[i][j] = rand.Float64()*20 - 10
        }
    }
    return m
}

func ValidateMatrix(m [][]float64) bool {
    n := len(m)
    if n == 0 {
        return false
    }
    for _, row := range m {
        if len(row) != n {
            return false
        }
    }
    return true
}

func DeterminantParallel(m [][]float64) float64 {
    n := len(m)
    if n == 1 {
        return m[0][0]
    }
    var det float64
    var wg sync.WaitGroup
    var mu sync.Mutex

    for col := 0; col < n; col++ {
        wg.Add(1)
        go func(col int) {
            defer wg.Done()
            sign := 1.0
            if col%2 != 0 {
                sign = -1.0
            }
            minor := getMinor(m, 0, col)
            d := Determinant(minor)
            mu.Lock()
            det += sign * m[0][col] * d
            mu.Unlock()
        }(col)
    }

    wg.Wait()
    return det
}

func Determinant(m [][]float64) float64 {
    n := len(m)
    if n == 1 {
        return m[0][0]
    }
    var det float64
    for col := 0; col < n; col++ {
        sign := 1.0
        if col%2 != 0 {
            sign = -1.0
        }
        minor := getMinor(m, 0, col)
        det += sign * m[0][col] * Determinant(minor)
    }
    return det
}

func getMinor(m [][]float64, row, col int) [][]float64 {
    minor := [][]float64{}
    for i := range m {
        if i == row {
            continue
        }
        newRow := []float64{}
        for j := range m[i] {
            if j == col {
                continue
            }
            newRow = append(newRow, m[i][j])
        }
        minor = append(minor, newRow)
    }
    return minor
}
