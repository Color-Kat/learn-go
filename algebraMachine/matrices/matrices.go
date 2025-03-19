package matrices

import (
	"algebraMachine/utils"
	"fmt"
	"strings"
)

/**
 * Matrix
 */

type Matrix[T int | bool] struct {
	Data [][]T
	Rows int
	Cols int
}

// NewMatrix creates a new matrix with the given rows, cols and neutral element
func NewMatrix[T int | bool](rows, cols int, neutralElement T) *Matrix[T] {
	matrix := &Matrix[T]{}

	matrix.Data = make([][]T, rows)
	matrix.Rows = rows
	matrix.Cols = cols

	for i := 0; i < rows; i++ {
		matrix.Data[i] = make([]T, cols)
		for j := 0; j < cols; j++ {
			matrix.Data[i][j] = neutralElement
		}
	}

	return matrix
}

// inputMatrix create a new matrix with the given rows, cols
// and populate it with the user input.
// Use parseString function to convert the input string to the desired type
func inputMatrix[T int | bool](cols, rows int, parseString func(s string) T) *Matrix[T] {
	matrix := &Matrix[T]{}

	matrix.Data = make([][]T, rows)
	matrix.Rows = rows
	matrix.Cols = cols

	for i := 0; i < rows; i++ {
	RepeatInput:
		input := strings.Fields(utils.Prompt(fmt.Sprintf("Insert row %v (%v values )", i, cols)))
		if len(input) != cols {
			utils.Prompt(fmt.Sprintf("Invalid value. Please enter %v values", cols))
			goto RepeatInput
		}

		matrix.Data[i] = make([]T, cols)
		matrix.Data[i] = utils.MapSlice(
			input,
			func(elem string) T {
				value := parseString(elem)
				return value
			},
		)
	}

	return matrix
}
