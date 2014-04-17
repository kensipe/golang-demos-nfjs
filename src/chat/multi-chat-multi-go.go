package main

import (
	"fmt"
	"net"
	"bytes"
)

type Caser interface {
	Case([]byte) []byte
}

type Uppercaser struct {
}

func (self Uppercaser) Case(input []byte) []byte {
	return bytes.ToUpper(input)
}

type Lowercaser struct {
}

func (self Lowercaser) Case(input []byte) []byte {
	return bytes.ToLower(input)
}

func main() {
	//	input()

	connection()
}

func connection() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	var clients []net.Conn
	input := make(chan []byte, 10)

	go chatManager(&clients, input, Uppercaser{})
	go chatManager(&clients, input, Lowercaser{})

	for {
		client, err := listener.Accept()
		if err != nil {
			continue
		}

		clients = append(clients, client)

		go handleClient(client, input)
	}

}

func chatManager(clients *[]net.Conn, input chan []byte, caser Caser) {
	for {
		message := <-input
		for _, client := range *clients {
			client.Write(caser.Case(message))
			prompt := []byte("> ")
			client.Write(prompt)
		}
	}
}

func handleClient(client net.Conn, input chan []byte) {

	for {
		buf := make([]byte, 4096)

		prompt := []byte("> ")
		client.Write(prompt)

		numBytes, err := client.Read(buf)
		if (numBytes == 0 || err != nil) {
			return
		}


		// add feature
		if (string(buf[:4]) == "exit") {
			client.Close()
			return
		}

		//		client.Write(buf)
		input <- buf
	}
}

func input() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

}
