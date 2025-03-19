package matrices

import "fmt"

type Operations struct {
	currentRow, currentCol, currentCol2 int
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
