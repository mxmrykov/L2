package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius    float64
	square    float64
	perimeter float64
}

type calc struct{}

func main() {
	rad := 3.0
	clc := calc{}

	square := clc.getSquare(rad)
	perimeter := clc.getPerimeter(rad)

	product := Circle{radius: rad, square: square, perimeter: perimeter}

	fmt.Println(product)
}

func (c calc) getSquare(rad float64) float64 {
	return math.Pi * math.Pow(rad, 2)
}

func (c calc) getPerimeter(rad float64) float64 {
	return 2 * math.Pi * rad
}
