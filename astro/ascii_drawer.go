package astro

func ScaleCoordinates(x float64, y float64, scale float64) (int, int) {
	return int(x * scale), int(y * scale)
}

func CreateGrid(size int) [][]rune {
	grid := make([][]rune, size)

	return grid
}
