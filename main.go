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
	iterations   = 90000
	inputLayer   = 2
	hiddenLayer  = 3
	outputLayer  = 1
)

func main() {
	inputMatrix, groundTruthLabels := data.GenerateData()
	NN := network.GenerateNetwork()
	NN = network.InitializeParameters(NN)

	var costs []float64
	var predictions [][][]float64

	//Train
	for epoch := 0; epoch < iterations; epoch++ {
		zMatrixOne, hMatrixOne, zMatrixTwo, prediction := network.Propagation(inputMatrix, NN)

		partialDerivativeOne, partialDerivativeTwo := network.Backpropagation(zMatrixOne, hMatrixOne, zMatrixTwo, prediction, groundTruthLabels, inputMatrix, NN)

		//Gradient Descent
		NN.WeightMatrixOne = maths.SubtractElementsMatrices(NN.WeightMatrixOne, maths.MultiplyValueToMatrices(learningRate, partialDerivativeOne))

		NN.WeightMatrixTwo = maths.SubtractElementsMatrices(NN.WeightMatrixTwo, maths.MultiplyValueToMatrices(learningRate, partialDerivativeTwo))

		epochCost := maths.CostBCE(prediction, groundTruthLabels)
		costs = append(costs, epochCost)
		predictions = append(predictions, prediction)
	}

	fmt.Println("---")
	fmt.Println("Training complete!")
	fmt.Printf("Last prediction: %v \n", predictions[len(predictions)-1])
	fmt.Printf("Expected labels: %v \n", groundTruthLabels)

	data.CreatePlot(costs, iterations)

	//Evaluate
	// fmt.Println(costs)
	// fmt.Println(predictions)
}
