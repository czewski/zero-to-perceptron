package main

import (
	"perceptron/pkg/biases"
)

func main() {
	//https://pub.towardsai.net/the-multilayer-perceptron-built-and-implemented-from-scratch-70d6b30f1964

	//Generate data

	//Generate NN
	input := biases.Layer{}
	hidden := biases.Layer{}
	output := biases.Layer{}

	NN := biases.Network{
		Inputs: []biases.Layer{input, input},
		Hidden: []biases.Layer{hidden, hidden, hidden},
		Output: []biases.Layer{output},
	}

	parametersInputHidden := make([][]float64, NN.LenInput())
	for i := 0; i < NN.LenInput(); i++ {
		parametersInputHidden[i] = make([]float64, NN.LenHidden())
	}

	//Train

	//Evaluate
}
