package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listerner, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error Listening:", err)
	}

	defer listerner.Close()

	for {
		conn, err := listerner.Accept()
		if err != nil {
			log.Fatal("Error accepting conn", conn)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	msg, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	ackMsg := strings.ToUpper(strings.TrimSpace(msg))
	response := fmt.Sprintf("ACK : %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server Write error: %v", err)
	}
}
