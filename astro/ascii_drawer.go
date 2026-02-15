package astro

import "fmt"

func ScaleCoordinates(x float64, y float64, scale float64) (int, int) {
	return int(x * scale), int(y * scale)
}

func CreateGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	return grid
}

func Plot(grid [][]rune, x, y int, symbol rune) {
	center := len(grid) / 2
	gridX := center + x
	gridY := center + y

	if gridX >= 0 && gridX < len(grid) && gridY >= 0 && gridY < len(grid) {
		grid[gridY][gridX] = symbol
	}
}

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
