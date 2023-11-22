package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
		if len(args) != 0 {
			if args[0] == ".." {
				p.route = routeUp(p.route)
			} else {
				w := []rune(args[0])
				if string(w[0]) == "/" {
					file, err := absPathErr(args[0])
					if err == nil {
						p.route = file
					} else {
						fmt.Println(err)
					}
				} else {
					file, err := tryGetPath(p.route, args[0])
					if err == nil {
						p.route = file
					} else {
						fmt.Println(err)
					}
				}
			}
		}
	case "pwd":
		p.logCurrentRoute()
	case "ps":
		logProcess()
	case "kill":
		Pid, errConv := strconv.Atoi(args[0])
		if errConv != nil {
			fmt.Println("Err at converting:", errConv)
			return
		}
		errKill := killProcById(Pid)
		if errKill != nil {
			fmt.Println("Err at killing:", errKill)
			return
		}
		fmt.Println("Process", args[0], "killed")
	case "ls":
		p.getAndLogEnv()
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

	builder := concString(newRoute)

	if len(strings.Split(builder, "\\")) <= 2 {
		return builder
	}

	return strings.TrimSuffix(builder, "\\")
}

func absPathErr(path string) (string, error) {
	if _, err := os.Stat(path); err == nil {
		pInfo, err := filepath.Abs(path)
		return pInfo, err
	} else {
		return "", err
	}
}

func tryGetPath(oldPath, path string) (string, error) {
	newPath := filepath.Join(oldPath, path)
	if _, err := os.Stat(newPath); err == nil {
		return newPath, nil
	} else {
		return "", err
	}
}

func concString(arr []string) string {

	var builder strings.Builder

	for _, s := range arr {
		_, err := builder.WriteString(s + "\\")
		if err != nil {
			log.Fatal(err)
		}
	}
	return builder.String()
}

func (p *proc) getAndLogEnv() {
	all, err := os.ReadDir(p.route)
	if err != nil {
		fmt.Println("Error at parsing enviroment:", err)
		return
	}
	for _, obj := range all {
		fmt.Println(obj)
	}
}

func logProcess() {
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println("Err at logging processes:", err)
		return
	}
	fmt.Println("   PID\t| Executable")
	for i := range processes {
		proc := processes[i]
		fmt.Printf("%d\t| %s\n", proc.Pid(), proc.Executable())
	}
}

func killProcById(Pid int) error {
	proc, errFind := os.FindProcess(Pid)
	if errFind != nil {
		return errFind
	}
	errKill := proc.Kill()
	if errFind != nil {
		return errKill
	}
	return nil
}

func (p *proc) logCurrentRoute() {
	fmt.Println()
	fmt.Println("Path")
	fmt.Println("----")
	fmt.Println(p.route)
	fmt.Println()
}
