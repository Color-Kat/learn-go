
package main

import (
	"algebraMachine/matrixOperations"
	//"algebraMachine/matrixOperations"
	"algebraMachine/matrixOperationsDeprecated"
	"algebraMachine/utils"
	"fmt"
	"strings"
)

func main() {
	matrix1 := matrixOperations.Matrix[int]{
		data: [][]int{
				{1, 2},
				{3, 4},
			},
		rows: 2,
		cols: 2
	}

	return

	matrix1 := matrixOperationsDeprecated.Matrix{}
	matrix2 := matrixOperationsDeprecated.Matrix{}

	for {
		utils.ClearConsole()
		option := printMenu()

		switch {
		case option == "1":
			inputMatrix(&matrix1)
		case option == "2":
			inputMatrix(&matrix2)
		case option == "3":
			utils.Prompt(fmt.Sprintln("MATRIX 1: \n", matrix1, "\nMatrix 2: \n", matrix2))
		case option == "4":
			chooseOperation(matrix1, matrix2)
		}

		_ = matrix1
		_ = matrix2
	}
}

func printMenu() string {
	fmt.Println("1. Set matrix 1")
	fmt.Println("2. Set matrix 2")
	fmt.Println("3. Print matrix")
	fmt.Println("4. Choose operation")
	fmt.Println("5. Exit")
	fmt.Println("--------------------")

	option := utils.Prompt("Choose an option: ")
	return option
}

func inputMatrix(matrix *matrixOperationsDeprecated.Matrix) {
	sizes := strings.Fields(utils.Prompt("Matrix size: "))

	if len(sizes) < 2 {
		utils.Prompt("Invalid input. Please enter two integers for rows and columns.")
		return
	}

	rows, err1 := utils.ToInt(sizes[0])
	cols, err2 := utils.ToInt(sizes[1])
	if err1 != nil || err2 != nil {
		utils.Prompt("Invalid input. Please enter two integers for rows and columns.")
		return
	}

	matrix.Populate(rows, cols)
}

func chooseOperation(matrixA, matrixB matrixOperationsDeprecated.Matrix) {
	operation := utils.Prompt(
		"Choose an operation: \n" +
			"1. MatrixAddition\n" +
			"2. Subtraction\n" +
			"3. Multiplication\n" +
			"4. Min\n" +
			"5. Max\n",
	)

	switch operation {
	case "1":
		matrixOperationsDeprecated.MatrixAddition(matrixA, matrixB)
	case "2":
		matrixOperationsDeprecated.MatrixSubtraction(matrixA, matrixB)
	default:
		utils.Prompt("Invalid operation")
	}
}
