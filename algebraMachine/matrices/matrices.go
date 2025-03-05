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

// InputIntMatrix create a new INT matrix with the given rows, cols
// and populate it with user input
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

// InputBoolMatrix create a new Bool matrix with the given rows, cols
// and populate it with user input
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
 * Operations
 */

type Operations struct {
	currentRow, currentCol, currentCol2 int
}

// IntMatrixOperation Specify the type of matrices for operations
type IntMatrixOperation struct {
	Operations
	MatrixA, MatrixB Matrix[int]
}

// BoolMatrixOperation Specify the type of matrices for operations
type BoolMatrixOperation struct {
	Operations
	MatrixA, MatrixB Matrix[bool]
}

type IMatrixOperations interface {
	Add() any
	Multiply() any
}

// Add Base function for addiction matrices.
// Define only the logic of the operation
func (mo *Operations) add(
	rowsCount1, colsCount1 int,
	rowsCount2, colsCount2 int,
	addFunc func(),
) {
	if rowsCount1 != rowsCount2 || colsCount1 != colsCount2 {
		fmt.Println("Rows and columns of matrices must be equal")
		return
	}

	// Change counters in the struct!
	for mo.currentRow = 0; mo.currentRow < rowsCount1; mo.currentRow++ {
		for mo.currentCol = 0; mo.currentCol < colsCount1; mo.currentCol++ {
			addFunc() // Call the function that defines the specific operation
		}
	}
}

// Multiply Base function for multiplication matrices
// Define only the logic of the operation
func (mo *Operations) multiply(
	rowsCount1, colsCount1 int,
	rowsCount2, colsCount2 int,
	multiplyFunc func(),
) {
	if colsCount1 != rowsCount2 {
		fmt.Println("Cols of matrix 1 must be equal to rows of matrix 2")
		return
	}

	// Change counters in the struct!
	for mo.currentRow = 0; mo.currentRow < rowsCount1; mo.currentRow++ {
		for mo.currentCol = 0; mo.currentCol < colsCount1; mo.currentCol++ {
			for mo.currentCol2 = 0; mo.currentCol2 < colsCount2; mo.currentCol2++ {
				multiplyFunc() // Call the function that defines the specific operation
			}
		}
	}
}

// Add sum 2 int matrices
func (mo *IntMatrixOperation) Add() any {
	// Create matrix with the result of operation
	resultMatrix := NewMatrix[int](mo.MatrixA.Rows, mo.MatrixA.Cols, 0)

	mo.add(
		mo.MatrixA.Rows, mo.MatrixA.Cols,
		mo.MatrixB.Rows, mo.MatrixB.Cols,
		func() {
			resultMatrix.Data[mo.currentRow][mo.currentCol] =
				mo.MatrixA.Data[mo.currentRow][mo.currentCol] +
					mo.MatrixB.Data[mo.currentRow][mo.currentCol]
		},
	)

	return resultMatrix
}

// Add sum 2 bool matrices
func (mo *BoolMatrixOperation) Add() any {
	// Create matrix with the result of operation
	resultMatrix := NewMatrix[bool](mo.MatrixA.Rows, mo.MatrixA.Cols, false)

	mo.add(
		mo.MatrixA.Rows, mo.MatrixA.Cols,
		mo.MatrixB.Rows, mo.MatrixB.Cols,
		func() {
			resultMatrix.Data[mo.currentRow][mo.currentCol] =
				mo.MatrixA.Data[mo.currentRow][mo.currentCol] ||
					mo.MatrixB.Data[mo.currentRow][mo.currentCol]
		},
	)

	return resultMatrix
}

// Multiply multiply 2 int matrices
func (mo *IntMatrixOperation) Multiply() any {
	// Create matrix with the result of operation
	// Rows of the first matrix and cols of the second matrix
	resultMatrix := NewMatrix[int](mo.MatrixA.Rows, mo.MatrixB.Cols, 0)

	mo.multiply(
		mo.MatrixA.Rows, mo.MatrixA.Cols,
		mo.MatrixB.Rows, mo.MatrixB.Cols,
		func() {
			//fmt.Println(
			//	"Calculation: ",
			//	resultMatrix.Data[mo.currentRow][mo.currentCol], "+=",
			//	mo.MatrixA.Data[mo.currentRow][mo.currentCol2], "*",
			//	mo.MatrixB.Data[mo.currentCol2][mo.currentCol], "=",
			//	resultMatrix.Data[mo.currentRow][mo.currentCol]+
			//		mo.MatrixA.Data[mo.currentRow][mo.currentCol2]*
			//			mo.MatrixB.Data[mo.currentCol2][mo.currentCol],
			//
			//	"Row:", mo.currentRow,
			//	"Col:", mo.currentCol,
			//	"Col2:", mo.currentCol2,
			//)

			resultMatrix.Data[mo.currentRow][mo.currentCol] +=
				mo.MatrixA.Data[mo.currentRow][mo.currentCol2] *
					mo.MatrixB.Data[mo.currentCol2][mo.currentCol]
		},
	)

	return resultMatrix
}

// Multiply multiply 2 int matrices
func (mo *BoolMatrixOperation) Multiply() any {
	// Create matrix with the result of operation
	// Rows of the first matrix and cols of the second matrix
	resultMatrix := NewMatrix[bool](mo.MatrixA.Rows, mo.MatrixB.Cols, false)

	mo.multiply(
		mo.MatrixA.Rows, mo.MatrixA.Cols,
		mo.MatrixB.Rows, mo.MatrixB.Cols,
		func() {
			resultMatrix.Data[mo.currentRow][mo.currentCol] =
				resultMatrix.Data[mo.currentRow][mo.currentCol] ||
					(mo.MatrixA.Data[mo.currentRow][mo.currentCol2] &&
						mo.MatrixB.Data[mo.currentCol2][mo.currentCol])
		},
	)

	return resultMatrix
}
