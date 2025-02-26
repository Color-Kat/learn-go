package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var toInt = strconv.Atoi

type Matrix struct {
	data [][]int
	rows int
	cols int
}

func (matrix *Matrix) set(cols, rows int) {
	matrix.data = make([][]int, rows)
	matrix.rows = rows
	matrix.cols = cols

	for i := 0; i < rows; i++ {
		matrix.data[i] = make([]int, cols)
		matrix.data[i] = mapSlice(
			strings.Fields(prompt(fmt.Sprintf("Insert row %v", i))),
			func(elem string) int {
				value, _ := toInt(elem)
				return value
			},
		)
		//for j := 0; j < cols; j++ {
		//	matrix.data[i][j] = 1
		//}
	}

}

func main() {

	matrix1 := Matrix{}
	matrix2 := Matrix{}

	for {
		clearConsole()
		option := printMenu()

		switch {
		case option == "1":
			inputMatrix(&matrix1)
		case option == "2":
			inputMatrix(&matrix2)
		case option == "3":
			prompt(fmt.Sprintln("MATRIX 1: \n", matrix1, "\nMatrix 2: \n", matrix2))
		case option == "4":
			chooseOperation(&matrix1, &matrix2)
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

	option := prompt("Choose an option: ")
	return option
}

func inputMatrix(matrix *Matrix) {
	sizes := strings.Fields(prompt("Matrix size: "))

	if len(sizes) < 2 {
		fmt.Println("Invalid input. Please enter two integers for rows and columns.")
		return
	}

	rows, err1 := toInt(sizes[0])
	cols, err2 := toInt(sizes[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please enter two integers for rows and columns.")
		return
	}

	matrix.set(rows, cols)
}

func chooseOperation(matrixA, matrixB *Matrix) {
	operation := prompt("Choose an operation: \n1. Addition\n2. Subtraction\n3. Multiplication\n4. Min\n5. Max")

	switch operation {
	case "1":
		addition(matrixA, matrixB)
	}
}

func addition(matrixA, matrixB *Matrix) {
	if matrixA.rows != matrixB.rows || matrixA.cols != matrixB.cols {
		fmt.Println("Matrices must have the same size")
		return
	}

	result := Matrix{}
	result.set(matrixA.rows, matrixA.cols)

	for i := 0; i < matrixA.rows; i++ {
		for j := 0; j < matrixA.cols; j++ {
			result.data[i][j] = matrixA.data[i][j] + matrixB.data[i][j]
		}
	}

	fmt.Println(result.data)
}

func prompt(prompt string) string {
	fmt.Println(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	return input
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func mapSlice[T any, U any](slice []T, callback func(elem T) U) []U {
	result := make([]U, len(slice))

	for i, value := range slice {
		result[i] = callback(value)
	}

	return result
}
