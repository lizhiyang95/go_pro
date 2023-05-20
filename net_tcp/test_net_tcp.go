package main

import (
	"fmt"
	"io"
	"net"
)

func tcpServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Listening on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		fmt.Println("Accepted connection from:", conn.RemoteAddr())
		io.WriteString(conn, "Hello, 世界!\n")
		conn.Close()
	}
}

func main() {
	tcpServer()
}
