package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {

	//получаем время с сервера ntp
	timeNtp, err := ntp.Time("pool.ntp.org")

	//парсим ошибки если они есть
	if err != nil {
		b, _ := fmt.Fprintf(os.Stderr, "Error at getting time: %v\n", err)
		fmt.Printf("%d bytes written", b)
		os.Exit(1)
	}

	//выводим время если все ок
	fmt.Printf("Current time: %s", timeNtp.Format(time.RFC3339))

}
