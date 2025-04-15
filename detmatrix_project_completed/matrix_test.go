package main

import "testing"

func TestDeterminant(t *testing.T) {
    matrix := [][]float64{
        {1, 2},
        {3, 4},
    }
    det := Determinant(matrix)
    if det != -2 {
        t.Errorf("Expected -2, got %f", det)
    }
}

func TestValidateMatrix(t *testing.T) {
    valid := [][]float64{{1, 2}, {3, 4}}
    invalid := [][]float64{{1, 2}, {3}}
    if !ValidateMatrix(valid) {
        t.Errorf("Matrix should be valid")
    }
    if ValidateMatrix(invalid) {
        t.Errorf("Matrix should be invalid")
    }
}
