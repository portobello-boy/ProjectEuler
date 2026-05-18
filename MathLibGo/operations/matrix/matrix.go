package matrix

import "fmt"

type Number interface {
	int | uint | float32 | float64
}

type Matrix[T Number] struct {
	Data [][]T
}

func NewMatrix[T Number](rows, cols int) *Matrix[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
		for j := range data[i] {
			data[i][j] = 0
		}
	}
	return &Matrix[T]{Data: data}
}

func (a Matrix[T]) Mult(b Matrix[T]) (*Matrix[T], error) {
	rowsA, colsA := len(a.Data), len(a.Data[0])
	rowsB, colsB := len(b.Data), len(b.Data[0])

	// Check if matrices can be multiplied
	if colsA != rowsB {
		return nil, fmt.Errorf("Matrix multiplication is not valid.")
	}

	// Initialize result matrix C
	C := NewMatrix[T](rowsA, colsB)

	// Perform multiplication
	for i := range rowsA {
		for j := range colsB {
			for k := range colsA {
				C.Data[i][j] += a.Data[i][k] * b.Data[k][j]
			}
		}
	}

	return &Matrix[T]{Data: C.Data}, nil
}
