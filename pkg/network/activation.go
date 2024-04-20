package network

import (
	"math"
)

func sigmoid(input float64) (activated float64) {
	activated = 1 / (1 + math.Pow(math.E, -input))
	return activated
}

func SigmoidMatrix(input [][]float64) [][]float64 {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			input[i][j] = sigmoid(input[i][j])
		}
	}

	return input
}

func SigmoidDerivative(input [][]float64) [][]float64 {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			input[i][j] = sigmoid(input[i][j]) * (1 - sigmoid(input[i][j]))
		}
	}

	return input
}
