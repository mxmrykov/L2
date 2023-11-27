package main

import "fmt"

type House struct {
	Walls      string
	Roof       string
	Windows    string
	Doors      string
	Foundation string
}

type HouseBuilder interface {
	BuildWalls() *HouseBuilder
	BuildRoof() *HouseBuilder
	BuildWindows() *HouseBuilder
	BuildDoors() *HouseBuilder
	BuildFoundation() *HouseBuilder
	GetHouse() *House
}

type ConcreteHouseBuilder struct {
	house *House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{house: &House{}}
}

func (b *ConcreteHouseBuilder) BuildWalls() *ConcreteHouseBuilder {
	b.house.Walls = "Brick"
	return b
}

func (b *ConcreteHouseBuilder) BuildRoof() *ConcreteHouseBuilder {
	b.house.Roof = "Tile"
	return b
}

func (b *ConcreteHouseBuilder) BuildWindows() *ConcreteHouseBuilder {
	b.house.Windows = "Glass"
	return b
}

func (b *ConcreteHouseBuilder) BuildDoors() *ConcreteHouseBuilder {
	b.house.Doors = "Wooden"
	return b
}

func (b *ConcreteHouseBuilder) BuildFoundation() *ConcreteHouseBuilder {
	b.house.Foundation = "Concrete"
	return b
}

func (b *ConcreteHouseBuilder) GetHouse() *House {
	return b.house
}

func main() {
	builder := NewConcreteHouseBuilder()
	house := builder.BuildWalls().BuildRoof().BuildWindows().BuildDoors().BuildFoundation().GetHouse()

	fmt.Println("Built house with the following features:")
	fmt.Printf("Walls: %s\n", house.Walls)
	fmt.Printf("Roof: %s\n", house.Roof)
	fmt.Printf("Windows: %s\n", house.Windows)
	fmt.Printf("Doors: %s\n", house.Doors)
	fmt.Printf("Foundation: %s\n", house.Foundation)
}
