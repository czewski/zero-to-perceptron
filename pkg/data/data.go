package data

func GenerateData() [][]float64 {
	designMatrix := [][]float64{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}

	return designMatrix
}

func GenerateLabels() [][]float64 {
	labels := [][]float64{
		{0}, {1}, {1}, {0},
	}

	return labels
}
