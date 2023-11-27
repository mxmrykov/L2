package main

import "fmt"

type Visitor interface {
	visitCircle(*Circle)
	visitSquare(*Square)
}

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitCircle(c)
}

type Square struct {
	sideLength float64
}

func (s *Square) accept(v Visitor) {
	v.visitSquare(s)
}

type AreaVisitor struct {
	totalArea float64
}

func (a *AreaVisitor) visitCircle(c *Circle) {
	area := 3.14 * c.radius * c.radius
	a.totalArea += area
}

func (a *AreaVisitor) visitSquare(s *Square) {
	area := s.sideLength * s.sideLength
	a.totalArea += area
}

func main() {
	shapes := []Shape{
		&Circle{radius: 5},
		&Square{sideLength: 4},
		&Circle{radius: 3},
	}

	areaVisitor := AreaVisitor{}

	for _, shape := range shapes {
		shape.accept(&areaVisitor)
	}

	fmt.Println("Total area of all shapes:", areaVisitor.totalArea)
}
