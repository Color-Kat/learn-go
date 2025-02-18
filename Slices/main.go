package main

import (
	"fmt"
)

func generateMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < cols; i++ {
		row := make([]int, 0)
		for j := 0; j < rows; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func main() {
	fmt.Print(generateMatrix(10, 10))
}
