//clock: TCP server for time

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer func() {
		fmt.Println("[Disconnect] 1 user left the game")
		c.Close()
	}()

	fmt.Println("[Connect] 1 user join the game.")
	for {
		_, err := io.WriteString(c, time.Now().Format("SB 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
