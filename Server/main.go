package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

const (
	connType = "tcp"
	connAddr = "127.0.0.1:8081"
)

// msgPacket
type msgPacket struct {
	Address string
	User    string
	Message string
}

func main() {
	fmt.Println("main: Initializing TCP Server...")
	ln, err := net.Listen(connType, connAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("main: Listening to " + connAddr)

	// create channel to receive messages from clients
	messageChan := make(chan msgPacket)
	go handleMessages(messageChan)

	// accepting new connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Println("main: ", conn.RemoteAddr(), " Connected")
		go handleConnection(conn, messageChan)
		defer conn.Close()
	}
}

// Handles all connections. Each client connection creates an instance of this goroutine
// (ex. if there are 7 clients, 7 handleConnection goroutines will exist)
func handleConnection(conn net.Conn, messageChan chan<- msgPacket) {
	for {
		err := gob.NewDecoder(conn).Decode(msg)
		if err != nil {
			log.Println("Nil")
			break
		}
		messageChan <- *msg
	}
}

// Handle Message Channel, receiving messages from all clients
func handleMessages(messageChan <-chan msgPacket) {
	for {
		msg := <-messageChan
		log.Println("handleMessages: received", msg)
	}
}
