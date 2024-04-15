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
	parametersInputHidden := make([][]float64, NN.LenInput())
	for i := 0; i < NN.LenInput(); i++ {
		parametersInputHidden[i] = make([]float64, NN.LenHidden())
	}

	//Parameters Hidden > Output (3x1) = 3
	parametersHiddenOutput := make([][]float64, NN.LenHidden())
	for i := 0; i < NN.LenInput(); i++ {
		parametersHiddenOutput[i] = make([]float64, NN.LenOutput())
	}

	//Initialize parameters with gaussian distribution

	//Z MATRIX: Multiply Input x InputHidden ---- row = each of the 4 training examples | column = z node
	zMatrix := maths.MultiplyMatrices(inputMatrix, parametersInputHidden)
	maths.PrintMatrix(zMatrix)

	//Activation function (at hidden layer) ---- apply sigmoid to each element in z matrix
	maths.PrintMatrix(network.SigmoidMatrix(zMatrix))

	//Get Z Matrix for Hidden Output

	//Activation for Hidden Output

	//Train

	//Evaluate
}
