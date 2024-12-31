package main

import (
	"fmt"
	"net"
	"log"
)

var port string = "8234"

func main() {
	// Запуск TCP-сервера
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to run a server: %v", err)
	}
	defer server.Close()
	fmt.Printf("Server was running on port: %s\n", port)

	// Бесконечный цикл принимающий подключения
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		// Обработка подключения в горутине
		go connectionHandler(connection)
	}
}

func connectionHandler(c net.Conn) {
	defer c.Close()

	// Чтение данных из подключения
	buffer := make([]byte, 1024)
	n, err := c.Read(buffer)
	if err != nil {
		fmt.Printf("Failed to read message from connection: %v\n", err)
		return
	}

	fmt.Printf("Message from connection(%s): %s\n", c.RemoteAddr().String(), string(buffer[:n]))


	// Отправка ответа
	_, err = c.Write([]byte("Hello from server!"))
	if err != nil {
		fmt.Printf("Faild to send message to connection: %v\n", err)
		return
	}
}