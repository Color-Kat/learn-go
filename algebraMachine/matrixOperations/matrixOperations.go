package matrixOperations

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

func NewMatrix[T int | bool](cols, rows int, neutralElement T) *Matrix[T] {
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

func inputMatrix[T int | bool](cols, rows int, parseString func(s string) T) *Matrix[T] {
	matrix := &Matrix[T]{}

	matrix.Data = make([][]T, rows)
	matrix.Rows = rows
	matrix.Cols = cols

	for i := 0; i < rows; i++ {
		matrix.Data[i] = make([]T, cols)
		matrix.Data[i] = utils.MapSlice(
			strings.Fields(utils.Prompt(fmt.Sprintf("Insert row %v (%v values )", i, cols))),
			func(elem string) T {
				value := parseString(elem)
				return value
			},
		)
	}

	return matrix
}

func InputIntMatrix(cols, rows int) *Matrix[int] {
	return inputMatrix[int](cols, rows, func(s string) int {
		value, err := utils.ToInt(s)
		if err != nil {
			utils.Prompt("Invalid input. Please enter an integer.")
			return 0
		}
		return value
	})
}

func InputBoolMatrix(cols, rows int) *Matrix[bool] {
	return inputMatrix[bool](cols, rows, func(s string) bool {
		value, err := utils.ToBool(s)
		if err != nil {
			utils.Prompt("Invalid input. Please enter a boolean.")
			return false
		}
		return value
	})
}

// ------------------------------

/**
 * MatrixOperation
 */

type MatrixOperation struct {
	currentRow, currentCol, currentCol2 int
}

// Add Base function for addiction matrices
func (mo *MatrixOperation) add(
	rowsCount1, colsCount1 int,
	rowsCount2, colsCount2 int,
	addFunc func(),
) {
	if rowsCount1 != rowsCount2 || colsCount1 != colsCount2 {
		panic("Rows and cols must be equal")
	}

	for mo.currentRow = 0; mo.currentRow < rowsCount1; mo.currentRow++ {
		for mo.currentCol = 0; mo.currentCol < colsCount1; mo.currentCol++ {
			addFunc()
		}
	}
}

// Multiply Base function for multiplication matrices
func (mo *MatrixOperation) multiply(
	rowsCount1, colsCount1 int,
	rowsCount2, colsCount2 int,
	multiplyFunc func(),
) {
	if colsCount1 != rowsCount2 {
		panic("Cols of matrix 1 must be equal to rows of matrix 2")
	}

	for mo.currentRow = 0; mo.currentRow < rowsCount1; mo.currentRow++ {
		for mo.currentCol2 = 0; mo.currentCol2 < colsCount2; mo.currentCol2++ {
			for mo.currentCol = 0; mo.currentCol < colsCount1; mo.currentCol++ {
				multiplyFunc()
			}
		}
	}
}

func (mo *MatrixOperation) AddInt(
	matrixA, matrixB Matrix[int],
) *Matrix[int] {
	resultMatrix := NewMatrix[int](matrixA.Cols, matrixA.Rows, 0)

	mo.add(
		matrixA.Rows, matrixA.Cols,
		matrixB.Rows, matrixB.Cols,
		func() {
			resultMatrix.Data[mo.currentRow][mo.currentCol] = matrixA.Data[mo.currentRow][mo.currentCol] + matrixB.Data[mo.currentRow][mo.currentCol]
		},
	)

	return resultMatrix
}

func (mo *MatrixOperation) AddBool(
	matrixA, matrixB Matrix[bool],
) *Matrix[bool] {
	resultMatrix := NewMatrix[bool](matrixA.Cols, matrixA.Rows, false)

	mo.add(
		matrixA.Rows, matrixA.Cols,
		matrixB.Rows, matrixB.Cols,
		func() {
			resultMatrix.Data[mo.currentRow][mo.currentCol] = matrixA.Data[mo.currentRow][mo.currentCol] || matrixB.Data[mo.currentRow][mo.currentCol]
		},
	)

	return resultMatrix
}
