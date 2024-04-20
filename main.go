package main

import (
	"perceptron/pkg/data"
	"perceptron/pkg/maths"
	"perceptron/pkg/network"
)

func main() {
	//https://pub.towardsai.net/the-multilayer-perceptron-built-and-implemented-from-scratch-70d6b30f1964

	//Generate data
	inputMatrix := data.GenerateData()
	groundTruthLabels := data.GenerateLabels()
	maths.PrintMatrix(inputMatrix)

	_, _ = inputMatrix, groundTruthLabels

	//Generate NN -- This still needs to make more sense
	input := network.Layer{}
	hidden := network.Layer{}
	output := network.Layer{}

	NN := network.Network{
		Inputs: []network.Layer{input, input},
		Hidden: []network.Layer{hidden, hidden, hidden},
		Output: []network.Layer{output},
	}

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

	//Initialize parameters with gaussian distribution //or random

	//Propagation
	//Z1 MATRIX: Multiply Input x InputHidden ---- row = each of the 4 training examples | column = z node
	zMatrixOne := maths.MultiplyMatrices(inputMatrix, weightMatrixOne)
	maths.PrintMatrix(zMatrixOne)

	//Activation function (at hidden layer) ---- apply sigmoid to each element in z matrix
	hMatrixOne := network.SigmoidMatrix(zMatrixOne)
	maths.PrintMatrix(hMatrixOne)

	//Get Z2 Matrix for Hidden Output

	zMatrixTwo := maths.MultiplyMatrices(hMatrixOne, weightMatrixTwo)
	maths.PrintMatrix(zMatrixTwo)

	//Activation for Hidden Output == Prediction
	prediction := network.SigmoidMatrix(zMatrixTwo)
	maths.PrintMatrix(prediction)

	//Backpropagation

	//Train

	//Evaluate
}
