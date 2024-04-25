package maths

import (
	"math"
)

func binaryCrossEntropy(predicted float64, groundTruth float64) (loss float64) {
	loss = -(groundTruth*math.Log(predicted) + (1-groundTruth)*math.Log(1-predicted))
	return loss
}

func CostBCE(predicted [][]float64, groundTruth [][]float64) (meanLoss float64) {
	summedLoss := 0.0

	rowsA, colsB := len(predicted), len(predicted[0])
	C := make([][]float64, rowsA)
	for i := range C {
		C[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			summedLoss += binaryCrossEntropy(predicted[i][j], groundTruth[i][j])
		}
	}

	// for indexTrainingExample := range predicted {
	// 	summedLoss += binaryCrossEntropy(predicted[indexTrainingExample], groundTruth[indexTrainingExample])
	// }

	meanLoss = summedLoss / float64(len(predicted))
	return meanLoss
}
