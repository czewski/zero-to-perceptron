package network

import (
	"perceptron/pkg/maths"
)

func Propagation(inputMatrix [][]float64, weightMatrixOne [][]float64, weightMatrixTwo [][]float64) (zMatrixOne, hMatrixOne, prediction [][]float64) {
	//Z1 MATRIX: Multiply Input x InputHidden ---- row = each of the 4 training examples | column = z node
	zMatrixOne = maths.MultiplyMatrices(inputMatrix, weightMatrixOne)

	//Activation function (at hidden layer) ---- apply sigmoid to each element in z matrix
	hMatrixOne = SigmoidMatrix(zMatrixOne)

	//Get Z2 Matrix for Hidden Output
	zMatrixTwo := maths.MultiplyMatrices(hMatrixOne, weightMatrixTwo)

	//Activation for Hidden Output == Prediction
	prediction = SigmoidMatrix(zMatrixTwo)

	return zMatrixOne, hMatrixOne, prediction
}
