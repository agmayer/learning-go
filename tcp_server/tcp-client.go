package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// Подключаемся к серверу
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// Чтение входных данных из stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем сообщение в сокет
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем порт
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}
