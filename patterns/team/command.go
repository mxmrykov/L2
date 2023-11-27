package main

import "fmt"

type Command interface {
	execute()
}

type Light struct{}

func (l *Light) turnOn() {
	fmt.Println("Light is turned on")
}

func (l *Light) turnOff() {
	fmt.Println("Light is turned off")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) execute() {
	c.light.turnOn()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) execute() {
	c.light.turnOff()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) pressButton() {
	r.command.execute()
}

func main() {
	light := Light{}

	lightOnCommand := LightOnCommand{&light}
	lightOffCommand := LightOffCommand{&light}

	remote := RemoteControl{}

	remote.command = &lightOnCommand
	remote.pressButton()

	remote.command = &lightOffCommand
	remote.pressButton()
}
