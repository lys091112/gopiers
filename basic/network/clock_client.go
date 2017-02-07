package network

import (
	"io"
	"log"
	"net"
	"os"
)

func Start_Clock_Client() {
	conn, err := net.Dial("tcp", "localhost:8099")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
