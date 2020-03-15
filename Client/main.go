package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	connAddr = "127.0.0.1"
	connPort = "8081"
)

func main() {
	fmt.Println("Connection to TCP Server...")
	conn, err := net.Dial("tcp", connAddr+":"+connPort)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		userMsg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// send userMsg to server
		fmt.Fprintf(conn, userMsg)
		serverMsg, err := bufio.NewReader(conn).ReadString('\n')
		serverMsg = strings.TrimSpace(serverMsg)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Resposta do servidor: " + serverMsg)
	}
}
