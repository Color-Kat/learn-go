package matrices

import "algebraMachine/utils"

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

// IntMatrixOperation Specify the type of matrices for operations
type IntMatrixOperation struct {
	Operations
	MatrixA, MatrixB Matrix[int]
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
