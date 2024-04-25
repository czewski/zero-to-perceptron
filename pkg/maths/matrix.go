package maths

import "fmt"

func DotMatrices(A, B [][]float64) [][]float64 {
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

func SubtractElementsMatrices(A, B [][]float64) [][]float64 {
	rowsA, colsB := len(A), len(B[0])

	//TODO: check for same dimension

	// Initialize result matrix C
	C := make([][]float64, rowsA)
	for i := range C {
		C[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}

	return C
}

func MultiplyElementWiseMatrices(A, B [][]float64) [][]float64 {
	rowsA, colsB := len(A), len(B[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
		for j := range result[i] {
			result[i][j] = A[i][0] * B[0][j]
		}
	}

	return result
}

func MultiplyValueToMatrices(A float64, B [][]float64) [][]float64 {
	rowsB, colsB := len(B), len(B[0])

	// Initialize result matrix C
	C := make([][]float64, rowsB)
	for i := range C {
		C[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsB; i++ {
		for j := 0; j < colsB; j++ {
			C[i][j] = A * B[i][j]
		}
	}

	return C
}

// https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func Transpose(slice [][]float64) [][]float64 {
	xl := len(slice[0])
	yl := len(slice)

	result := make([][]float64, xl)
	for i := range result {
		result[i] = make([]float64, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func PrintMatrix(matrix [][]float64) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
