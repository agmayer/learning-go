package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// Слушаем порт tcp localhost:8081
	serverLine, _ := net.Listen("tcp", ":8081")

	// Открываем порт
	openConnect, _ := serverLine.Accept()

	// Запускаем цикл
	for {
		// Будем слушать все сообщения разделённые \n
		message, _ := bufio.NewReader(openConnect).ReadString('\n')
		// Распечатываем полученное сообщение
		fmt.Print("Message Recived:", string(message))
		// Процесс выборки для полученной строки
		newMessage := strings.ToUpper(message)
		// Отправить новую строку обратно клиенту
		openConnect.Write([]byte(newMessage + "\n"))
	}
}
