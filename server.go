package main

import (
	"fmt"
	"main/handler"
	"main/types"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	handler.ErrorHandler(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handler.ErrorHandler(err)

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	for {
		payload, err := types.Decode(conn)
		handler.ErrorHandler(err)

		fmt.Println("Client's message: ", string(payload.Bytes()))

		var response types.Binary
		response = types.Binary("Server received your message: " + string(payload.Bytes()))

		_, err = response.WriteTo(conn)
		conn.SetDeadline(time.Now().Add(time.Second * 5))
		if err != nil {
			typeErr, ok := err.(net.Error)
			if typeErr.Timeout() && ok {
				fmt.Println("Write timeout!")
			} else {
				handler.ErrorHandler(typeErr)
			}
		}
	}
}
