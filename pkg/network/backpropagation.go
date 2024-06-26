package network

import (
	"perceptron/pkg/maths"
)

func Backpropagation(zMatrixOne, hMatrixOne, zMatrixTwo, prediction, groundTruthLabels, inputMatrix [][]float64, NN Network) (partialDerivativeOne, partialDerivativeTwo [][]float64) {
	//∂loss/∂p *∂p/∂zh * ∂zh/∂wh
	delOne := maths.SubtractElementsMatrices(prediction, groundTruthLabels)

	partialDerivativeTwo = maths.DotMatrices(maths.Transpose(hMatrixOne), delOne)

	//∂loss/∂p * ∂p/∂zh * ∂zh/∂h * ∂h/∂z * ∂z/∂w
	delTwo := maths.MultiplyElementWiseMatrices(delOne, maths.Transpose(NN.WeightMatrixTwo))

	delThree := maths.MultiplyElementWiseMatrices(delTwo, SigmoidDerivative(zMatrixOne))

	partialDerivativeOne = maths.DotMatrices(maths.Transpose(inputMatrix), delThree)

	return partialDerivativeOne, partialDerivativeTwo
}
