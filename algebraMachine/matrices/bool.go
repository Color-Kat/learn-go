package matrices

import "algebraMachine/utils"

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

// BoolMatrixOperation Specify the type of matrices for operations
type BoolMatrixOperation struct {
	Operations
	MatrixA, MatrixB Matrix[bool]
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

// Multiply multiply 2 bool matrices
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
