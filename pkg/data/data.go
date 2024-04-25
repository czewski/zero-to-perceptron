package data

func GenerateData() (designMatrix, labels [][]float64) {
	return generateInput(), generateLabels()
}

func generateInput() [][]float64 {
	designMatrix := [][]float64{
		{1, 0},
		{0, 1},
		{0, 0},
		{1, 1},
	}

	return designMatrix
}

func generateLabels() [][]float64 {
	labels := [][]float64{
		{1}, {1}, {0}, {0},
	}

	return labels
}
