package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

		for i := 0; i < len(prm); i += 1 {
			if i%2 == 0 {
				inp.params = append(inp.params, prm[i:i+2])
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
	_, errRead := ReadFile(mainPath)
	if errRead != nil {
		return
	}

	fmt.Println(generateRegex(I.parseFlagsContext()["-a"], I.parseFlagsContext()["-b"], I.word))
}

func generateRegex(a, b int, word string) string {
	res := fmt.Sprintf("^[\\w\\d ]{0,%d}(%s)[\\w\\d ]{0,%d}$", a, word, b)

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
