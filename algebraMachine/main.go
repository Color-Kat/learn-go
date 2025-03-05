package main

import (
	"algebraMachine/matrixOperations"
	"algebraMachine/utils"
	"fmt"
	"strings"
)

var ring string
var matrices = map[string][]any{
	"int":  {nil, nil},
	"bool": {nil, nil},
}

func main() {
Main:
	for {
		utils.ClearConsole()
		option := printMenu()

		switch {
		case option == "1":
			chooseRing()
		case option == "2":
			inputMatrix(0)
		case option == "3":
			inputMatrix(1)
		case option == "4":
			utils.Prompt(fmt.Sprintln("MATRIX 1: \n", matrices[ring][0], "\nMatrix 2: \n", matrices[ring][1]))
			//case option == "5":
			//	chooseOperation(matrix1, matrix2)
		default:
			break Main
		}

	}
}

func printMenu() string {
	fmt.Println("1. Choose ring")
	fmt.Println("2. Set matrix 1")
	fmt.Println("3. Set matrix 2")
	fmt.Println("4. Print matrix")
	fmt.Println("5. Choose operation")
	fmt.Println("6. Exit")
	fmt.Println("--------------------")

	option := utils.Prompt("Choose an option: ")
	return option
}

func chooseRing() {
	fmt.Println("__Choose ring__")
	fmt.Println("1. Integers")
	fmt.Println("2. Booleans")

	switch utils.Prompt("Enter option: ") {
	case "1":
		ring = "int"
	case "2":
		ring = "bool"
	}
}

func inputMatrix(matrixIndex int) {
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

	switch ring {
	case "int":
		matrices[ring][matrixIndex] = matrixOperations.InputIntMatrix(cols, rows)
	case "bool":
		matrices[ring][matrixIndex] = matrixOperations.InputBoolMatrix(cols, rows)
	default:
		utils.Prompt("Invalid ring")
	}
}

//func chooseOperation[T int | bool](matrixA, matrixB matrixOperations.IMatrix) {
//operation := utils.Prompt(
//	"Choose an operation: \n" +
//		"1. MatrixAddition\n" +
//		"2. Subtraction\n" +
//		"3. Multiplication\n" +
//		"4. Min\n" +
//		"5. Max\n",
//)
//
//switch operation {
//case "1":
//	matrixOperations.MatrixAddition(matrixA, matrixB)
//case "2":
//	matrixOperations.MatrixSubtraction(matrixA, matrixB)
//default:
//	utils.Prompt("Invalid operation")
//}
//}
