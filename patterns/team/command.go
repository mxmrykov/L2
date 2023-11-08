package main

import "fmt"

func (c *Controller) execute() {
	fmt.Println("Command executed")
}

type Controller struct{}

type Receiver struct {
	c *Controller
}

func main() {

	printContr := &Controller{}

	receiver := Receiver{printContr}

	receiver.c.execute()
}
