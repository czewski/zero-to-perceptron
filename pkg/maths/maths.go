package maths

import "fmt"

func MultiplyMatrices(A, B [][]float64) [][]float64 {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])

	// Check if matrices can be multiplied
	if colsA != rowsB {
		fmt.Println("Matrix multiplication is not valid.")
		return nil
	}

	// Initialize result matrix C
	C := make([][]float64, rowsA)
	for i := range C {
		C[i] = make([]float64, colsB)
	}

	// Perform multiplication
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

func PrintMatrix(matrix [][]float64) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
