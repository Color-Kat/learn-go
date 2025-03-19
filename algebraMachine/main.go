package main

import (
	"algebraMachine/floydWarshall"
	"algebraMachine/matrices"
	"algebraMachine/utils"
	"fmt"
	"math"
	"strings"
)

var testIntMatrix = matrices.Matrix[int]{
	Rows: 2,
	Cols: 2,
	Data: [][]int{
		{1, 2},
		{3, 4},
	},
}

// Answer for Square
//  7 10
// 15 22

var testBoolMatrix = matrices.Matrix[bool]{
	Rows: 2,
	Cols: 2,
	Data: [][]bool{
		{true, false},
		{true, true},
	},
}

// Answer for Square
// true false
// true true

var ring = "int" // Current ring for operations
var matricesByRing = map[string][]any{
	"int":  {testIntMatrix, testIntMatrix},
	"bool": {testBoolMatrix, testBoolMatrix},
}

func main() {
Main:
	for {
		option := printMenu()

		switch {
		case option == "1":
			chooseRing()
		case option == "2":
			inputMatrix(0)
		case option == "3":
			inputMatrix(1)
		case option == "4":
			utils.ClearConsole()
			utils.Prompt(fmt.Sprintln("MATRIX 1: \n", matricesByRing[ring][0], "\nMatrix 2: \n", matricesByRing[ring][1]))
		case option == "5":
			chooseOperation()
		default:
			break Main
		}
	}
}

func printMenu() string {
	utils.ClearConsole()
	fmt.Println("____ Matrix operation ____")
	fmt.Println("__ on an arbitrary ring __")
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
	utils.ClearConsole()
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
	utils.ClearConsole()
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
		matricesByRing[ring][matrixIndex] = *matrices.InputIntMatrix(cols, rows)
	case "bool":
		matricesByRing[ring][matrixIndex] = *matrices.InputBoolMatrix(cols, rows)
	default:
		utils.Prompt("Invalid ring")
	}
}

func chooseOperation() {
	utils.ClearConsole()
	operation := utils.Prompt(
		"Choose an operation: \n" +
			"1. Addition\n" +
			//"2. Subtraction\n" +
			"2. Multiplication\n" +
			"3. Floyd Warshall (only for int)\n",
		//"3. Min\n" +
		//"4. Max\n",
	)

	// Create instance of struct with specific ring operations
	// It's easy to implement more rings
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
		utils.Prompt(fmt.Sprint(operations.Add()))
	case "2":
		utils.Prompt(fmt.Sprint(operations.Multiply()))
	case "3":
		inf := math.MaxInt32
		graph := [][]int{
			{0, 4, 11},
			{6, 0, 2},
			{3, inf, 0},
		}

		//graph = matricesByRing["int"][0].(matrices.Matrix[int]).Data

		fmt.Println("___FLoyd Warshall algorithm___")
		fmt.Println("Initial graph:")
		floydWarshall.PrintFloydWarshallGraph(graph)

		result, next := floydWarshall.FloydWarshall(graph)

		fmt.Println("Graph of the shortest distances:")
		floydWarshall.PrintFloydWarshallGraph(result)

		start, end := 0, 2
		path := floydWarshall.FindPath(next, start, end)

		if path != nil {
			fmt.Printf("Path from %d to %d: %v\n", start, end, path)
			fmt.Printf("Distance: %d\n", result[start][end])
		} else {
			fmt.Printf("No path from %d to %d\n", start, end)
		}

		utils.Prompt("Press any key")
	default:
		utils.Prompt("Invalid operation")
	}
}
