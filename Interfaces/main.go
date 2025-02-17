package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (rect rect) area() float64 {
	return rect.width * rect.height
}

func (rect rect) perimeter() float64 {
	return rect.width*2 + rect.height*2
}

type circle struct {
	radius float64
}

func (circle circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (circle circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func main() {
	myRect := rect{
		width:  10,
		height: 5,
	}

	myCircle := circle{
		radius: 10,
	}

	describeShape(myRect)
	describeShape(myCircle)
}

func describeShape(shape shape) {
	fmt.Printf("This shape has area equals to %v\n", shape.area())
	fmt.Printf("And perimeter equals to %v\n", shape.perimeter())
}
