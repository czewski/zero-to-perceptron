package main

import (
	"perceptron/pkg/data"
	"perceptron/pkg/maths"
	"perceptron/pkg/network"
)

// ANN definitions
const (
	learningRate = 0.1
	iterations   = 100000
	inputLayer   = 2
	hiddenLayer  = 3
	outputLayer  = 1
)

func main() {
	// Init
	inputMatrix, groundTruthLabels := data.GenerateData()
	NN := network.GenerateNetwork()

	//Initialize Parameters
	weightMatrixOne, weightMatrixTwo := network.InitializeParameters(NN)
	var cost [][]float64
	var predictions [][][]float64

	//Train
	for epoch := 0; epoch < iterations; epoch++ {
		zMatrixOne, hMatrixOne, prediction := network.Propagation(inputMatrix, weightMatrixOne, weightMatrixTwo)
		partialDerivativeOne, partialDerivativeTwo := network.Backpropagation(zMatrixOne, hMatrixOne, prediction, groundTruthLabels, inputMatrix, weightMatrixTwo)

		//Gradient Descent
		weightMatrixOne = maths.SubtractElementsMatrices(weightMatrixOne, maths.MultiplyValueToMatrices(learningRate, partialDerivativeOne))
		weightMatrixTwo = maths.SubtractElementsMatrices(weightMatrixOne, maths.MultiplyValueToMatrices(learningRate, partialDerivativeTwo))

		//Calculate cost -- CostBCE will return a slice of means (?)
		epochCost := maths.CostBCE(prediction, groundTruthLabels)
		cost = append(cost, epochCost)

		predictions = append(predictions, prediction)
	}

	//Evaluate
}
