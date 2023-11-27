package main

import "fmt"

type CPU struct{}

func (c *CPU) start() {
	fmt.Println("Starting CPU")
}

func (c *CPU) execute() {
	fmt.Println("Executing CPU instructions")
}

type Memory struct{}

func (m *Memory) load() {
	fmt.Println("Loading data into memory")
}

type HardDrive struct{}

func (h *HardDrive) read() {
	fmt.Println("Reading data from hard drive")
}

type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) start() {
	cf.cpu.start()
	cf.memory.load()
	cf.hardDrive.read()
	cf.cpu.execute()
}

func main() {
	computer := NewComputerFacade()
	computer.start()
}
