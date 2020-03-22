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

type loginPacket struct {
	Address string
	User    string
}

type messagePacket struct {
	SourceUser      string
	DestinationUser string
	Message         string
}

// UsersHash = username -> conn
var UsersHash = make(map[string]net.Conn)

func main() {
	fmt.Println("main: Initializing TCP Server...")
	ln, err := net.Listen(connType, connAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("main: Listening to " + connAddr)

	// create channel to receive messages from clients
	messageChan := make(chan messagePacket)
	go handleMessages(messageChan)

	// accepting new connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		login := &loginPacket{}
		err = gob.NewDecoder(conn).Decode(login)
		if err != nil {
			log.Println("main: Error decoding login message from", conn.RemoteAddr())
			break
		}
		UsersHash[login.User] = conn
		log.Println("main: New user registered:", login.User, login.Address)
		fmt.Println("main:", UsersHash)
		go handleConnection(conn, messageChan, login.User)
	}
}

// Handles all connections. Each client connection creates an instance of this goroutine
// (ex. if there are 7 clients, 7 handleConnection goroutines will exist)
func handleConnection(conn net.Conn, messageChan chan<- messagePacket, user string) {
	for {
		msg := &messagePacket{}
		err := gob.NewDecoder(conn).Decode(msg)
		if err != nil {
			log.Println("handleConnection: Connection closed with", user)
			delete(UsersHash, user)
			break
		}
		messageChan <- *msg
	}
	defer conn.Close()
}

// Handle Message Channel, receiving messages from all clients
func handleMessages(messageChan <-chan messagePacket) {
	for {
		msg := <-messageChan
		destConn := UsersHash[msg.DestinationUser]
		log.Println("handleMessages: received", msg)
		err := gob.NewEncoder(destConn).Encode(msg)
		if err != nil {
			log.Println("handleMessages: Error encoding message", err)
		}
	}
}
