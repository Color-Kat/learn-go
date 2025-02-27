package matrixOperations

import (
	"algebraMachine/utils"
	"fmt"
	"strings"
)

type Matrix struct {
	data [][]int
	rows int
	cols int
}

func (matrix *Matrix) Init(cols, rows int, neutralElement int) {
	matrix.data = make([][]int, rows)
	matrix.rows = rows
	matrix.cols = cols

	for i := 0; i < rows; i++ {
		matrix.data[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix.data[i][j] = neutralElement
		}
	}
}

func (matrix *Matrix) Populate(cols, rows int) {
	matrix.data = make([][]int, rows)
	matrix.rows = rows
	matrix.cols = cols

	for i := 0; i < rows; i++ {
		matrix.data[i] = make([]int, cols)
		matrix.data[i] = utils.MapSlice(
			strings.Fields(utils.Prompt(fmt.Sprintf("Insert row %v (%v values )", i, cols))),
			func(elem string) int {
				value, _ := utils.ToInt(elem)
				return value
			},
		)
	}
}

func MatrixAddition(matrixA, matrixB Matrix) {
	if matrixA.rows != matrixB.rows || matrixA.cols != matrixB.cols {
		utils.Prompt("Matrices must have the same size")
		return
	}

	result := Matrix{}
	result.Init(matrixA.rows, matrixA.cols, 0)

	for i := 0; i < matrixA.rows; i++ {
		for j := 0; j < matrixA.cols; j++ {
			result.data[i][j] = matrixA.data[i][j] + matrixB.data[i][j]
		}
	}

	utils.Prompt(fmt.Sprintln("Result:\n", result.data))
}

func MatrixSubtraction(matrixA, matrixB Matrix) {
	// Inverse elements of matrix b
	matrixB.data = utils.MapSlice(matrixB.data, func(row []int) []int {
		return utils.MapSlice(row, func(elem int) int {
			return -elem
		})
	})

	MatrixAddition(matrixA, matrixB)
}

func MatrixMultiplication(matrixA, matrixB *Matrix) {
	utils.Prompt("Not implemented yet")
}

func MatrixMin(matrixA, matrixB *Matrix) {
	utils.Prompt("Not implemented yet")
}

func MatrixMax(matrixA, matrixB *Matrix) {
	utils.Prompt("Not implemented yet")
}
