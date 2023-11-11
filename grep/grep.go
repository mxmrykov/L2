package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type input struct {
	command, path, word string
	params              [][]string
}

func main() {

	//создаем буфер прослушки вводимых в консоль даныных
	reader := bufio.NewReader(os.Stdin)

	//логаем изначальный заголовок
	fmt.Println("Grep Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("command> ")

		//слушаем строку из буфера, деленной по \n
		buf, _ := reader.ReadString('\n')

		splited := strings.Split(buf, " ")

		command := splited[0]
		command = strings.TrimSuffix(command, "\n")
		command = strings.TrimSuffix(command, "\r")

		//завершаем программу если команда exit
		if command == "exit" {
			os.Exit(0)
		}

		//проверяем что введена и команда и второй аргумент, флаги опциональны
		if len(splited) < 3 {
			fmt.Println("Wrong arguments")
			continue
		}

		inp := &input{}

		inp.command = command

		inp.path = splited[len(splited)-1]
		inp.path = strings.TrimSuffix(inp.path, "\n")
		inp.path = strings.TrimSuffix(inp.path, "\r")

		inp.params = make([][]string, 0)
		inp.word = splited[len(splited)-2]

		prm := splited[1 : len(splited)-2]

		flagMustSkip := false

		for i := 0; i < len(prm); i += 1 {
			if prm[i] != "-i" && prm[i] != "-v" && prm[i] != "-n" {
				if flagMustSkip {
					flagMustSkip = false
				} else {
					inp.params = append(inp.params, prm[i:i+2])
					flagMustSkip = true
				}
			} else {
				inp.params = append(inp.params, prm[i:i+1])
			}
		}

		inp.searchForWord()

	}

}

func (I *input) searchForWord() {

	//проверяем если команда не сорт - значит неизвестная
	if I.command != "grep" {
		fmt.Println("Unknown command")
		return
	}

	//для наилучешго взаимодействия проинициализируем глобальный путь к исполнаяемому файлу
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	//склеиваем путь из консоли и глобальный путь
	mainPath := filepath.Join(path, I.path)

	//получаем список наших слов
	words, errRead := ReadFile(mainPath)
	if errRead != nil {
		return
	}

	bothRegs := false

	for _, flag := range I.params {
		if flag[0] == "-i" {
			bothRegs = true
		}
	}

	searchRegex := generateRegex(I.word, bothRegs)

	regex, _ := regexp.Compile(searchRegex)

	searched := [][]string{}

	for _, word := range words {
		temp := regex.FindAllString(word, -1)
		searched = append(searched, temp)
	}

	var builder strings.Builder

	for _, wds := range searched {
		for _, wd := range wds {
			_, er := builder.WriteString(strings.TrimPrefix(wd+"\n", " "))
			if er != nil {
				fmt.Println(er)
			}
		}
	}

	splitedReady := strings.Split(builder.String(), "\n")
	parsedContext := I.parseFlagsContext()

	for _, elem := range splitedReady {
		tempElem := strings.Trim(elem, "\n\r")
		for _, el := range getIndexOf(words, tempElem) {
			if el != 0 {
				if len(words[el-parsedContext["-b"]]) != 0 {
					fmt.Println(words[el-parsedContext["-b"]])
				}
				fmt.Println(words[el])
				if len(words[el+parsedContext["-a"]]) != 0 {
					fmt.Println(words[el-parsedContext["-a"]])
				}
			}
		}
	}
}

func getIndexOf(ar []string, elem string) []int {

	indexes := []int{0}

	for i := 0; i < len(ar); i += 1 {
		if ar[i] == elem {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func generateRegex(word string, bothRegs bool) string {

	res := fmt.Sprintf(".*\\b(%s)\\b.*", word)

	if bothRegs {
		res = fmt.Sprintf(".*\\b(%s|%s)\\b.*", strings.ToUpper(word), strings.ToLower(word))
	}

	return res
}

func (I *input) parseFlagsContext() map[string]int {

	ret := make(map[string]int, 3)

	ret["-a"] = 0
	ret["-b"] = 0
	ret["-C"] = 0

	for _, flag := range I.params {
		if flag[0] == "-a" || flag[0] == "-b" || flag[0] == "-C" {
			ret[flag[0]], _ = strconv.Atoi(flag[1])
		}
	}

	return ret

}

// чтение файла и возврат данных
func ReadFile(mainPath string) ([]string, error) {

	//открываем файл по указанному пути
	cont, errParse := os.Open(mainPath)
	if errParse != nil {
		fmt.Printf("Error at opening file: %v\n", errParse)
		return nil, errParse
	}

	//читаем данные
	scan := bufio.NewScanner(cont)

	//заранее создаем массив слов взятых из файла
	var words []string

	//добавляем содержимое файла в массив слов построчно
	for scan.Scan() {
		words = append(words, scan.Text())
	}

	//если все ок возвращаем массив слов и нулевую ошибку
	return words, nil
}
