package main

import (
	"algebraMachine/matrices"
	"algebraMachine/utils"
	"fmt"
	"strings"
)

var ring string = "int"
var matricesByRing = map[string][]any{
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
			utils.Prompt(fmt.Sprintln("MATRIX 1: \n", matricesByRing[ring][0], "\nMatrix 2: \n", matricesByRing[ring][1]))
		case option == "5":
			chooseOperation()
		default:
			break Main
		}

	}
}

func printMenu() string {
	fmt.Println("1. Choose ring")
	fmt.Println("2. Set matrix 1")
	fmt.Println("3. Set matrix 2")
	fmt.Println("4. Print matrices")
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
		matricesByRing[ring][matrixIndex] = matrices.InputIntMatrix(cols, rows)
	case "bool":
		matricesByRing[ring][matrixIndex] = matrices.InputBoolMatrix(cols, rows)
	default:
		utils.Prompt("Invalid ring")
	}
}

func chooseOperation() {
	operation := utils.Prompt(
		"Choose an operation: \n" +
			"1. Addition\n" +
			"2. Subtraction\n" +
			"3. Multiplication\n" +
			"4. Min\n" +
			"5. Max\n",
	)

	var operations matrices.IMatrixOperations
	switch ring {
	case "int":
		operations = &matrices.IntMatrixOperation{
			MatrixA: matricesByRing[ring][0].(matrices.Matrix[int]),
			MatrixB: matricesByRing[ring][1].(matrices.Matrix[int]),
		}
	case "bool":
		operations = &matrices.BoolMatrixOperation{
			MatrixA: matricesByRing[ring][0].(matrices.Matrix[bool]),
			MatrixB: matricesByRing[ring][1].(matrices.Matrix[bool]),
		}
	}

	switch operation {
	case "1":
		fmt.Println(operations.Add())
	case "2":
		//matrices.MatrixSubtraction(matrixA, matrixB)
	default:
		utils.Prompt("Invalid operation")
	}
}
