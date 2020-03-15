package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	connType = "tcp"
	connAddr = "127.0.0.1"
	connPort = "8081"
)

func main() {
	fmt.Println("main: Initializing TCP Server...")
	ln, err := net.Listen(connType, connAddr+":"+connPort)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("main: Listening to " + connAddr + " : " + connPort)

	// create channel to receive messages from clients
	messageChan := make(chan string)
	go handleMessages(messageChan)

	// accepting new connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Println("main: ", conn.RemoteAddr(), " Connected")
		go handleConnection(conn, messageChan)
	}
}

func handleConnection(conn net.Conn, messageChan chan string) {
	for {
		// read connection message
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		msg = strings.TrimSpace(msg)

		// check if client requested disconnection
		temp := strings.ToUpper(msg)
		if temp == "STOP" {
			log.Println("handleConnection: ", conn.RemoteAddr(), " requested disconnection")
			break
		}

		// answer connection message and sendo to messages channel
		conn.Write([]byte("Mensagem recebida \n"))
		messageChan <- string(conn.RemoteAddr().String()) + ": " + msg
	}
	conn.Close()
}

func handleMessages(messageChan chan string) {
	for {
		msg, ok := <-messageChan
		log.Println("handleMessage: received message from messageChan: "+msg+" (chan:", ok, ")")
	}
}
