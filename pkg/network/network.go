package network

// ixj i = inputs, j = hidden
type Network struct {
	Inputs []Layer
	Hidden []Layer
	Output []Layer
}

// Neuron ou layer que tem a activation? neuron i guess
type Layer struct {
	Neurons []Neuron
	Bias    float64
}

type Neuron struct {
	Input           []float64
	PreviousOutputs []float64
	Output          float64
	//Bias            Bias
	// Activation
}

// type Bias struct {
// 	Value float64
// }

type Activation struct {
}

func (a Activation) Sigmoid([][]float64) (as []float64) {

	return as
}

func (nn Network) LenInput() int {
	return len(nn.Inputs)
}
func (nn Network) LenHidden() int {
	return len(nn.Hidden)
}
func (nn Network) LenOutput() int {
	return len(nn.Output)
}

//Network : Input, Hidden, Output
// 2,3,1
// 6, 3
//Create [6][3]float64
//
