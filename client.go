package main

import (
	"fmt"
	"main/handler"
	"main/types"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:1234")
	handler.ErrorHandler(err)
	defer dial.Close()
	message := "Hello"
	data := types.Binary(message)
	_, err = data.WriteTo(dial)
	handler.ErrorHandler(err)
	dial.SetDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		typeErr, ok := err.(net.Error)
		if typeErr.Timeout() && ok {
			fmt.Println("Server timeout!")
		} else {
			handler.ErrorHandler(typeErr)
		}
	}
	payload, err := types.Decode(dial)
	fmt.Println("Response Server: ", string(payload.Bytes()))
}
