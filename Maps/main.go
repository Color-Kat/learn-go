package main

import (
	"fmt"
)

type Coordinates struct {
	x int
	y int
}

func (cords *Coordinates) run(xOffset, yOffset int) {
	cords.x += xOffset
	cords.y += yOffset
}

// -----------------

type NestedKey struct {
	Country, City string
}

func main() {
	// First way to create a map
	map1 := make(map[string]int)
	map1["ColorKat"] = 100
	map1["X"] = 50

	fmt.Println(map1)

	// Nested map in go
	positions := make(map[string]map[string]Coordinates)
	positions["ColorKat"] = make(map[string]Coordinates)
	positions["ColorKat"]["queen"] = Coordinates{
		x: 0,
		y: 0,
	}
	positions["ColorKat"]["king"] = Coordinates{
		x: 1,
		y: 1,
	}

	fmt.Println(positions)

	// Move queen
	position := positions["ColorKat"]["queen"]
	position.run(8, 0)
	positions["ColorKat"]["queen"] = position

	// Move king
	position = positions["ColorKat"]["king"]
	position.run(0, 1)
	positions["ColorKat"]["king"] = position

	// Move pawn (initialize with x: 0, y: 0)
	position, ok := positions["ColorKat"]["pawn"]
	position.run(0, 1)
	positions["ColorKat"]["pawn"] = position
	if !ok {
		fmt.Println("Pawn is added to the board")
	}

	fmt.Println(positions)

	// Nested map with typed keys
	todoList := make(map[NestedKey]string)
	todoList[NestedKey{"USA", "Los Angeles"}] = "Me"

	fmt.Println(todoList[NestedKey{"USA", "Los Angeles"}])

}
