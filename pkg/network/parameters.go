package network

func InitializeParameters(NN Network) ([][]float64, [][]float64) {

	//Parameters Input > Hidden (2x3) = 6
	weightMatrixOne := make([][]float64, NN.LenInput())
	for i := 0; i < NN.LenInput(); i++ {
		weightMatrixOne[i] = make([]float64, NN.LenHidden())
	}

	//Parameters Hidden > Output (3x1) = 3 // 3x1 or 1x3 to make it work?
	weightMatrixTwo := make([][]float64, NN.LenHidden())
	for i := 0; i < NN.LenHidden(); i++ {
		weightMatrixTwo[i] = make([]float64, NN.LenOutput())
	}

	//TODO: Initialize parameters with gaussian distribution //or random

	return weightMatrixOne, weightMatrixTwo
}
