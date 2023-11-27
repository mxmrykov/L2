package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("Drawing a circle")
}

type Rectangle struct{}

func (r *Rectangle) draw() {
	fmt.Println("Drawing a rectangle")
}

type ShapeFactory struct{}

func (sf *ShapeFactory) createShape(shapeType string) Shape {
	if shapeType == "circle" {
		return &Circle{}
	} else if shapeType == "rectangle" {
		return &Rectangle{}
	} else {
		return nil
	}
}

func main() {
	factory := ShapeFactory{}

	circle := factory.createShape("circle")
	circle.draw()

	rectangle := factory.createShape("rectangle")
	rectangle.draw()
}
