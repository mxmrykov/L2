package main

import (
	"fmt"
	"log"
)

type Logger struct {
	logger  Log
	printer Print
}

type Log struct{}
type Print struct{}

func main() {
	logger := Logger{}

	logger.printer.Print("message printed from subclass Print")
	logger.logger.Log("message printed from subclass Log")
}

func (l Log) Log(message string) {
	log.Println(message)
}

func (p Print) Print(message string) {
	fmt.Println(message)
}
