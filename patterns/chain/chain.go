package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) // метод для установки следующего обработчика
	Handle(request string)   // метод для обработки запроса
}

type ConcreteHandlerA struct {
	next Handler
}

func (c *ConcreteHandlerA) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA: Handling request A")
	} else if c.next != nil {
		fmt.Println("ConcreteHandlerA: Passing request to next handler")
		c.next.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerA: Cannot handle request")
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (c *ConcreteHandlerB) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB: Handling request B")
	} else if c.next != nil {
		fmt.Println("ConcreteHandlerB: Passing request to next handler")
		c.next.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerB: Cannot handle request")
	}
}

func main() {
	handlerA := ConcreteHandlerA{}
	handlerB := ConcreteHandlerB{}

	handlerA.SetNext(&handlerB)

	handlerA.Handle("A")
	handlerA.Handle("B")
	handlerA.Handle("C")
}
