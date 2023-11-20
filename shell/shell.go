package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type proc struct {
	route string
}

func main() {
	//создаем буфер прослушки вводимых в консоль даныных
	reader := bufio.NewReader(os.Stdin)

	//логаем изначальный заголовок
	fmt.Println("Unix Shell")
	fmt.Println("---------------------")

	dir, errCurPath := os.Getwd()

	if errCurPath != nil {
		fmt.Printf("Cannot resolve path: %v\n", errCurPath)
	}

	main := &proc{route: dir}

	//слушаем консоль в бесконечном цикле
	for {
		fmt.Printf("Shell | %s> ", main.route)

		//слушаем строку из буфера, деленной по \n
		buf, _ := reader.ReadString('\n')

		//отделяем нулевой элемент вводимой строки от всех отсальных, это будет команда,
		//в нашем случае либо sort либо exit. Также триммируем суффиксы для получения
		//конечной команды в виде строки
		command := strings.Split(buf, " ")[0]
		command = strings.TrimSuffix(command, "\n")
		command = strings.TrimSuffix(command, "\r")

		//завершаем программу если команда \quit
		if command == "\\quit" {
			os.Exit(0)
		}

		//проверяем что введена и команда
		if len(strings.Split(buf, " ")) < 1 {
			fmt.Println("Wrong arguments")
			continue
		} else if len(strings.Split(buf, " ")) == 1 {
			main.parseCommand(command, nil)
		} else {
			args := strings.Split(buf, " ")[1:]

			args[len(args)-1] = strings.TrimSuffix(args[len(args)-1], "\n")
			args[len(args)-1] = strings.TrimSuffix(args[len(args)-1], "\r")

			main.parseCommand(command, args)
		}
	}
}

func (p *proc) parseCommand(command string, args []string) {
	switch command {
	case "echo":
		printClearArray(args)
	case "cd":
		if args[0] == ".." {
			p.route = routeUp(p.route)
		}
	case "pwd":
		p.logCurrentRoute()
	}
}

func printClearArray(args []string) {
	for _, i := range args {
		fmt.Print(i + " ")
	}
	fmt.Println()
}

func routeUp(currentRoute string) string {
	splitedRoute := strings.Split(currentRoute, "\\")
	newRoute := splitedRoute[:len(splitedRoute)-1]

	var builder strings.Builder
	for _, s := range newRoute {
		_, err := builder.WriteString(s + "\\")
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(strings.Split(builder.String(), "\\")) <= 2 {
		return builder.String()
	}

	return strings.TrimSuffix(builder.String(), "\\")
}

func (p *proc) logCurrentRoute() {
	fmt.Println()
	fmt.Println("Path")
	fmt.Println("----")
	fmt.Println(p.route)
	fmt.Println()
}
