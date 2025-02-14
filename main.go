package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("___ BMI Calculator ___\n\n")

	height, weight := getSizes()

	bmi := calculateBMI(height, weight)

	printResult(bmi)
}

func getSizes() (float64, float64) {
	var height float64
	var weight float64

	fmt.Println("Enter your height in cm:")
	fmt.Scan(&height)
	fmt.Println("Enter your weight in kg:")
	fmt.Scan(&weight)

	return height, weight
}

func calculateBMI(height float64, weight float64) float64 {
	return weight / math.Pow(height/100, 2)
}

func printResult(mbi float64) {
	fmt.Printf("Your BMI is %.0f\n", mbi)
	fmt.Println("You should lose a little weight.")
}
