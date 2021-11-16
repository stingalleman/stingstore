package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	store map[string]string = make(map[string]string)
)

func CreateServer(host string, port string) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Printf("Listening for connections on %s:%s\n", host, port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	msg := fmt.Sprintf("%s", buf[:n])

	if msg == "list" {
		for i, v := range store {
			conn.Write([]byte(fmt.Sprintf("%s=%s", i, v)))
		}
	} else if strings.Contains(msg, "=") {
		msg := strings.Split(msg, "=")
		key := msg[0]
		val := msg[1]

		store[key] = val
	} else {
		conn.Write([]byte(store[msg]))
	}

	conn.Close()
}
