package main

import (
	"bufio"
	"fmt"
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

	cont, errParse := os.OpenFile(mainPath, os.O_RDWR|os.O_TRUNC, 0755)

	if errParse != nil {
		fmt.Printf("Error at opening file: %v\n", errParse)
		return
	}

	defer cont.Close()
	scan := bufio.NewScanner(cont)

	var words []string

	for scan.Scan() {
		words = append(words, scan.Text())
	}

	if len(inp.flags) == 0 {
		var (
			justWrite       = make([]string, len(words))
			finalJustString = make([][]byte, 0, len(words))
		)
		copy(justWrite, words)
		sort.Strings(justWrite)
		for _, i := range justWrite {
			finalJustString = append(finalJustString, []byte(i+"\n"))
		}
		fmt.Println(finalJustString)
		for _, i := range finalJustString {
			_, erWrite := cont.Write(i)
			if erWrite != nil {
				fmt.Printf("Error at wtiting string %s to file: %v\n", string(i), erWrite)
			}
		}
		fmt.Println("file sorted")
		return
	}

	switch inp.flags {
	}
	return
}

//jh
//vb
//nv
//bn
//vd
//as
//fd
//gdg
//dfg
//f
