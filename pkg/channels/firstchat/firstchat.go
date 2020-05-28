package main

import "net"

func main() {

	// creates chat server which we can telnet to
	// telnet localhost 9000
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	for {
		client, err := listener.Accept()
		if err != nil {
			continue
		}

		handleNewClient(client)
		// now try it with a go
		//go handleNewClient(client)
	}
}

func handleNewClient(client net.Conn) {

	for {
		buf := make([]byte, 4096)

		numBytes, err := client.Read(buf)
		if numBytes == 0 || err != nil {
			return
		}

		// add feature
		if string(buf[:4]) == "exit" {
			return
		}

		client.Write(buf)
	}
}
