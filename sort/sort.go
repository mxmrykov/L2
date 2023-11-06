package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
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

		if strings.Split(buf, "\n")[0] == "exit" {
			os.Exit(0)
		}

		if len(strings.Split(buf, " ")) < 2 {
			fmt.Println("Wrong arguments")
			continue
		}

		returnable := &input{}

		returnable.command = strings.Split(buf, " ")[0]
		returnable.path = strings.Split(buf, " ")[1]
		returnable.flags = strings.Split(buf, " ")[2:]

		returnable.parseInput()
	}
}
func (inp *input) parseInput() {

	if inp.command != "sort" {
		fmt.Println("Unknown command")
		return
	}

	path, er := os.Getwd()

	if er != nil {
		log.Println(er)
	}

	fmt.Println(filepath.Join(path, inp.path))

	cont, errParse := ioutil.ReadFile(filepath.Join(path, inp.path))

	if errParse != nil {
		fmt.Printf("Error at opening file: %v\n", errParse)
		return
	}

	switch inp.flags {
	case nil:
		wordsDefault := strings.Split(string(cont), "\n")
		sort.Strings(wordsDefault)
		for _, i := range wordsDefault {
			errWrite := ioutil.WriteFile(filepath.Join(path, inp.path), []byte(i+"\n"), 0666)

			if errWrite != nil {
				fmt.Printf("Error at writing file: %v\n", errWrite)
			}
		}
	}
	return
}
