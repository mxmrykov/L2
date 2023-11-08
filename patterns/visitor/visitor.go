package main

import (
	"fmt"
	"math/rand"
)

type Visitor struct {
	name  string
	genID int
}

type Gen struct{}

func main() {

	gen := &Gen{}

	maxim := &Visitor{name: "maxim"}
	ivan := &Visitor{name: "ivan"}

	gen.genID(maxim)
	gen.genID(ivan)

	fmt.Println(*maxim)
	fmt.Println(*ivan)
}

func (g *Gen) genID(vis *Visitor) {
	vis.genID = 1000 + rand.Intn(8999)
}
