package network

import (
	"math"
	"math/rand"
)

func InitializeParameters(NN Network) Network { //([][]float64, [][]float64)
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

	NN.WeightMatrixOne = initializeRandomValues(weightMatrixOne)
	NN.WeightMatrixTwo = initializeRandomValues(weightMatrixTwo)

	return NN
}

func initializeRandomValues(slice [][]float64) [][]float64 {
	for i := range slice {
		for j := range slice[i] {
			slice[i][j] = rand.NormFloat64() * math.Sqrt(2.0/float64(len(slice[i])))

			//slice[i][j] = rand.NormFloat64() * math.Sqrt(1.0/float64(len(slice[i])))
		}
	}
	return slice
}
