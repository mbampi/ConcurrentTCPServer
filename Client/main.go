package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	connAddr = "127.0.0.1:8081"
)

// msgPacket
type msgPacket struct {
	Address string
	User    string
	Message string
}

func main() {
	fmt.Println("Connection to TCP Server...")

	conn, err := net.Dial("tcp", connAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		log.Println(err)
	}
	msg := &msgPacket{
		Address: conn.LocalAddr().String(),
		User:    username,
	}

	for {
		// Get input from user
		fmt.Print("Message: ")
		reader = bufio.NewReader(os.Stdin)
		userMsg, err := reader.ReadString('\n')
		userMsg = strings.TrimSpace(userMsg)
		if err != nil {
			log.Fatal(err)
		}

		// Send to server
		msg.Message = userMsg
		gob.NewEncoder(conn).Encode(msg)
		log.Println("Sent to server:", msg)
	}
}
