package biases

// ixj i = inputs, j = hidden
type Network struct {
	Inputs []Layer
	Hidden []Layer
	Output []Layer
}

// Neuron ou layer que tem a activation? neuron i guess
type Layer struct {
	Neurons []Neuron
	Bias    float32
}

type Neuron struct {
	Input           []float32
	PreviousOutputs []float32
	Output          float32
	//Bias            Bias
	// Activation
}

// type Bias struct {
// 	Value float32
// }

type Activation struct {
}

func (a Activation) Sigmoid([][]float32) (as []float32) {

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
//Create [6][3]float32
//
