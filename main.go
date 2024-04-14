package main

import (
	"perceptron/pkg/biases"
	"perceptron/pkg/data"
)

func main() {
	//https://pub.towardsai.net/the-multilayer-perceptron-built-and-implemented-from-scratch-70d6b30f1964

	//Generate data
	inputMatrix := data.GenerateData()
	groundTruthLabels := data.GenerateLabels()
	_, _ = inputMatrix, groundTruthLabels

	//Generate NN -- This needs to make more sense still
	input := biases.Layer{}
	hidden := biases.Layer{}
	output := biases.Layer{}

	NN := biases.Network{
		Inputs: []biases.Layer{input, input},
		Hidden: []biases.Layer{hidden, hidden, hidden},
		Output: []biases.Layer{output},
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

	//Linear transformation from input to hidden

	//Train

	//Evaluate
}
