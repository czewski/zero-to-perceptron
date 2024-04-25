package network

import "math/rand"

func InitializeParameters(NN Network) ([][]float64, [][]float64) {
	//Parameters Input > Hidden (2x3) = 6
	weightMatrixOne := make([][]float64, NN.LenInput())
	for i := 0; i < NN.LenInput(); i++ {
		weightMatrixOne[i] = make([]float64, NN.LenHidden())
	}

	//Parameters Hidden > Output (3x1) = 3
	weightMatrixTwo := make([][]float64, NN.LenHidden())
	for i := 0; i < NN.LenHidden(); i++ {
		weightMatrixTwo[i] = make([]float64, NN.LenOutput())
	}

	//Initialize random values
	weightMatrixOne = initializeRandomValues(weightMatrixOne)
	weightMatrixTwo = initializeRandomValues(weightMatrixTwo)

	return weightMatrixOne, weightMatrixTwo
}

func initializeRandomValues(slice [][]float64) [][]float64 {
	for i := range slice {
		for j := range slice[i] {
			slice[i][j] = rand.NormFloat64()
		}
	}
	return slice
}
