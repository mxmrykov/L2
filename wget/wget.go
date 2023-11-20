package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type object struct {
	filename, href string
	wFileName      bool
	params         [][]string
	limit          int
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

		//завершаем программу если команда exit
		if command == "exit" {
			os.Exit(0)
		}

		//проверяем что введена и команда и второй аргумент, флаги опциональны
		if len(splited) < 2 {
			fmt.Println("Wrong arguments")
			continue
		}

		href := splited[len(splited)-1]
		href = strings.TrimSuffix(href, "\n")
		href = strings.TrimSuffix(href, "\r")

		params := make([][]string, 0)

		prm := splited[1 : len(splited)-2]

		flagMustSkip := false

		for i := 0; i < len(prm); i += 1 {
			if flagMustSkip {
				flagMustSkip = false
			} else {
				params = append(params, prm[i:i+2])
				flagMustSkip = true
			}
		}

		ob := object{
			href:   href,
			params: params,
		}
		err := ob.download()

		if err != nil {
			fmt.Println(err)
		}

	}
}

func (o *object) download() error {

	file, errGet := http.Get(o.href)
	if errGet != nil {
		fmt.Println("Error at getting file: ", errGet)
		return errGet
	}

	for _, flag := range o.params {
		if flag[0] == "-O" {
			o.wFileName = true
			o.filename = flag[1]
		}
		if strings.Split(flag[0], "=")[0] == "––limit-rate" {
			o.limit, _ = strconv.Atoi(strings.Split(flag[0], "=")[1])
		}
	}

	defer file.Body.Close()

	var downloadedBytes int

	buf := make([]byte, 1024)
	prevTime := time.Now()
	prevDownloadedBytes := 0

	for {
		n, err := file.Body.Read(buf)
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(prevTime).Seconds()
		downloadSpeed := float64(downloadedBytes-prevDownloadedBytes) / 1024 / elapsedTime
		currentString := [100]byte{}
		percentsFloat := float64(downloadedBytes) / float64(file.ContentLength) * 100
		percents := int(percentsFloat)
		for i := 0; i < percents; i += 1 {
			currentString[i] = '#'
		}
		for i := percents; i < 100; i += 1 {
			currentString[i] = '.'
		}
		if n > 0 {
			downloadedBytes += n
			fmt.Printf("%d of %d KB downloaded |%s| %d %s | Speed : %.2f KB/s\r", downloadedBytes/1024, file.ContentLength/1024, currentString, percents, "%", downloadSpeed)
		}
		if err != nil {
			break
		}
	}

	fmt.Println("Download successful\nFinal file size:", downloadedBytes/1024, "KB")

	defer file.Body.Close()

	if !o.wFileName {
		o.filename = path.Base(file.Request.URL.String())
	}

	cr, errCreate := os.Create(o.filename)
	if errCreate != nil {
		fmt.Println("Error at creating file: ", errCreate)
		return errCreate
	}

	defer cr.Close()

	_, errWrite := io.Copy(cr, file.Body)

	if errWrite != nil {
		fmt.Println("Error at writing at file: ", errWrite)
		return errWrite
	}

	return nil
}
