package main

import (
	"fmt"
	"net"
	"log"
)

var addr string = "localhost"
var port string = "8234"

func main() {
	// Подключение к серверу
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()

	// Отправка сообщения на сервер
	message := "Hello from client"
	_, err = connection.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message to server: %v", err)
	}
	fmt.Printf("Sended message: %s\n", message)

	// Чтение ответа от сервера
	buffer := make([]byte, 1024)
	n, err := connection.Read(buffer)
	if err != nil {
		log.Fatalf("Failed to read response message: %v", err)
	}

	fmt.Printf("Response message: %s\n", string(buffer[:n]))
}