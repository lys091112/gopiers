package network

import (
	"io"
	"log"
	"net"
	"time"
)
// StartClock test
func StartClock() {
	listener, err := net.Listen("tcp", "localhost:8099")
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

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("14:02:04\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
