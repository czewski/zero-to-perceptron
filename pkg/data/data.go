package data

func GenerateData() [][]float32 {
	designMatrix := [][]float32{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}

	return designMatrix
}

func GenerateLabels() []float32 {
	labels := []float32{
		0, 1, 1, 0,
	}

	return labels
}
