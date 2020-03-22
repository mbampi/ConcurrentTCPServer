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
	connType = "tcp"
	connAddr = "127.0.0.1:8081"
)

type loginPacket struct {
	Address string
	User    string
}

type messagePacket struct {
	SourceUser      string
	DestinationUser string
	Message         string
}

func receiveMessages(server net.Conn) {
	for {
		msg := &messagePacket{}
		err := gob.NewDecoder(server).Decode(msg)
		if err != nil {
			log.Println("receiveMessages: Error decoding received message", err)
		}
		fmt.Println(msg.SourceUser+":", msg.Message)
	}
}

func main() {
	fmt.Println("Connection to TCP Server...")

	conn, err := net.Dial(connType, connAddr)
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
	login := &loginPacket{
		User:    username,
		Address: conn.LocalAddr().String(),
	}
	gob.NewEncoder(conn).Encode(login)
	go receiveMessages(conn)

	msgPacket := &messagePacket{
		SourceUser: username,
	}
	for {
		bufio.NewReader(os.Stdin).ReadString('\n')
		// Ask for destination user
		fmt.Print("To: ")
		destinationUser, err := bufio.NewReader(os.Stdin).ReadString('\n')
		destinationUser = strings.TrimSpace(destinationUser)
		if err != nil {
			log.Fatal(err)
		}

		// Ask for input message
		fmt.Print("Message: ")
		msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
		msg = strings.TrimSpace(msg)
		if err != nil {
			log.Fatal(err)
		}

		// Send to server
		msgPacket.DestinationUser = destinationUser
		msgPacket.Message = msg
		gob.NewEncoder(conn).Encode(msgPacket)
		log.Println("Sent to server:", msgPacket)
	}
}
