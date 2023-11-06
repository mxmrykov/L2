package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type input struct {
	command, path string
	flags         []string
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Sort Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("command> ")
		buf, _ := reader.ReadString('\n')

		command := strings.Split(buf, " ")[0]

		command = strings.TrimSuffix(command, "\n")
		command = strings.TrimSuffix(command, "\r")

		if command == "exit" {
			os.Exit(0)
		}

		if len(strings.Split(buf, " ")) < 2 {
			fmt.Println("Wrong arguments")
			continue
		}

		returnable := &input{}

		returnable.command = command
		returnable.path = strings.Split(buf, " ")[1]

		for _, i := range strings.Split(buf, " ")[2:] {
			i = strings.TrimSuffix(i, "\n")
			i = strings.TrimSuffix(i, "\r")
			returnable.flags = append(returnable.flags, i)
		}

		returnable.path = strings.TrimSuffix(returnable.path, "\n")
		returnable.path = strings.TrimSuffix(returnable.path, "\r")

		returnable.parseInput()
	}
}
func (inp *input) parseInput() {

	if inp.command != "sort" {
		fmt.Println("Unknown command")
		return
	}

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	mainPath := filepath.Join(path, inp.path)

	cont, errParse := os.ReadFile(mainPath)

	if errParse != nil {
		fmt.Printf("Error at opening file: %v\n", errParse)
		return
	}

	//var words []string
	//
	//for _, i :=

	fmt.Println(strings.ReplaceAll(string(cont), "\n", " "))

	switch inp.flags {
	case nil:
		//wordsDefault := strings.Split(string(cont), "\n")
		//sort.Strings(wordsDefault)
		//fmt.Println(wordsDefault)
		//for _, i := range wordsDefault {
		//	errWrite := os.WriteFile(mainPath, []byte(i+"\n"), 0666)
		//	if errWrite != nil {
		//		fmt.Printf("Error at writing file: %v\n", errWrite)
		//	}
		//}
	}
	return
}
