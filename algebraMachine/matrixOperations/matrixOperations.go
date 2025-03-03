package matrixOperations

import (
	"algebraMachine/utils"
	"fmt"
	"strings"
)

type Matrix[T int | bool] struct {
	data [][]T
	rows int
	cols int
}

func (matrix *Matrix[T]) Init(cols, rows int, neutralElement T) {
	matrix.data = make([][]T, rows)
	matrix.rows = rows
	matrix.cols = cols

	for i := 0; i < rows; i++ {
		matrix.data[i] = make([]T, cols)
		for j := 0; j < cols; j++ {
			matrix.data[i][j] = neutralElement
		}
	}
}

func (matrix *Matrix[T]) Populate(cols, rows int, parseString func(s string) T) {
	matrix.data = make([][]T, rows)
	matrix.rows = rows
	matrix.cols = cols

	for i := 0; i < rows; i++ {
		matrix.data[i] = make([]T, cols)
		matrix.data[i] = utils.MapSlice(
			strings.Fields(utils.Prompt(fmt.Sprintf("Insert row %v (%v values )", i, cols))),
			func(elem string) T {
				value := parseString(elem)
				return value
			},
		)
	}
}

type MatrixOperation struct {
	currentRow, currentCol, currentCol2 int
}

// Add Base function for addiction matrices
func (mo *MatrixOperation) Add(
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

func (mo *MatrixOperation) Multiply(
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
) {
	resultMatrix := make([][]int, matrixA.rows)
	mo.Add(
		matrixA.rows, matrixA.cols,
		matrixB.rows, matrixB.cols,
		func() {
			resultMatrix[mo.currentRow][mo.currentCol] = matrixA.data[mo.currentRow][mo.currentCol] + matrixB.data[mo.currentRow][mo.currentCol]
		},
	)
}
