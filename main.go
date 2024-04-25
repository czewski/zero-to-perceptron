package main

import (
	"fmt"
	"perceptron/pkg/data"
	"perceptron/pkg/maths"
	"perceptron/pkg/network"
)

// ANN definitions
const (
	learningRate = 0.1
	iterations   = 10000000
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
	var costs []float64
	var predictions [][][]float64
	var zMatrixOne, hMatrixOne, prediction, partialDerivativeOne, partialDerivativeTwo [][]float64

	//Train
	for epoch := 0; epoch < iterations; epoch++ {
		zMatrixOne, hMatrixOne, prediction = network.Propagation(inputMatrix, weightMatrixOne, weightMatrixTwo)

		partialDerivativeOne, partialDerivativeTwo = network.Backpropagation(zMatrixOne, hMatrixOne, prediction, groundTruthLabels, inputMatrix, weightMatrixTwo)

		//Gradient Descent
		weightMatrixOne = maths.SubtractElementsMatrices(weightMatrixOne, maths.MultiplyValueToMatrices(learningRate, partialDerivativeOne))

		weightMatrixTwo = maths.SubtractElementsMatrices(weightMatrixTwo, maths.MultiplyValueToMatrices(learningRate, partialDerivativeTwo))

		epochCost := maths.CostBCE(prediction, groundTruthLabels)

		costs = append(costs, epochCost)
		predictions = append(predictions, prediction)
	}

	fmt.Print(len(predictions))

	fmt.Println("---")
	fmt.Println("Training complete!")
	fmt.Printf("Last prediction: %v \n", predictions[len(predictions)-1])
	fmt.Printf("Expected labels: %v \n", groundTruthLabels)

	data.CreatePlot(costs, iterations)

	//Evaluate
	// fmt.Println(costs)
	// fmt.Println(predictions)
}
