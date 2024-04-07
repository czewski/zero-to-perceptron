package biases

type Network struct {
	Full []Layer
}

// Neuron ou layer que tem a activation? neuron i guess
type Layer struct {
	Neurons  []Neuron
	Bias     float32
	IsInput  bool
	IsOutput bool
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
