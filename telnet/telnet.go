package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type connection struct {
	address, port string
	socket        net.Conn
	params        []string
}

func main() {
	//создаем буфер прослушки вводимых в консоль даныных
	reader := bufio.NewReader(os.Stdin)

	//логаем изначальный заголовок
	fmt.Println("Telnet Shell")
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
		if len(splited) < 3 {
			fmt.Println("Wrong arguments")
			continue
		}

		address := splited[1]

		port := splited[2]

		port = strings.TrimSuffix(port, "\n")
		port = strings.TrimSuffix(port, "\r")

		var tempFlags []string

		//парсим флаги по аналогии с командой, триммируя \n и \r
		for _, i := range strings.Split(buf, " ")[3:] {
			i = strings.TrimSuffix(i, "\n")
			i = strings.TrimSuffix(i, "\r")
			tempFlags = append(tempFlags, i)
		}

		conn := &connection{address: address, port: port, params: tempFlags}

		conn.start()
	}
}

func (c connection) start() {
	address := fmt.Sprintf("%s:%s", c.address, c.port)
	fmt.Println("Trying", address, "...")

	timeOut := 0 * time.Second

	if len(c.params[0]) != 0 {
		prvTime := strings.TrimSuffix(strings.Split(c.params[0], "=")[1], "s")
		prvTimeInt, _ := strconv.Atoi(prvTime)
		timeOut = time.Duration(prvTimeInt) * time.Second
	}

	conn, errConnect := net.DialTimeout("tcp", address, timeOut)

	if errConnect != nil {
		fmt.Println("Err at read stdin:", errConnect)
		os.Exit(1)
	}

	c.socket = conn

	defer c.socket.Close()
	fmt.Println("Connected to", address)

	go read(c)

	go listen(c)
}

func read(c connection) {
	buf := make([]byte, 1024)
	for {
		inp, err := c.socket.Read(buf)

		if err != nil {
			fmt.Println("Err at read stdin:", err)
			os.Exit(1)
		}

		fmt.Println(inp)
	}
}

func listen(c connection) {
	defer c.socket.Close()
	buf := make([]byte, 1024)
	for {
		fmt.Print("telnet> ")
		inp, errRead := os.Stdin.Read(buf)

		if errRead != nil {
			fmt.Println("Err at read stdin:", errRead)
			os.Exit(1)
		}

		_, errSockWrite := c.socket.Write(buf[:inp])

		if errSockWrite != nil {
			fmt.Println("Err at write socket:", errSockWrite)
			os.Exit(1)
		}

	}
}
