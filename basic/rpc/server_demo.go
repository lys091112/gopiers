package rpc

import (
	"log"
	"net"
	"net/rpc"
)

const FirstServiceName = "/base/rpc.HelloService"

type FristInterface = interface {
	Hello(request string, response *string) error
}

func RegisterFirstService(svc FristInterface) error {
	return rpc.RegisterName(FirstServiceName, svc)
}

type HelloService struct{}

func (s *HelloService) Hello(request string, response *string) error {
	*response = "hello:" + request
	return nil
}

func Start() {
	RegisterFirstService(new(HelloService))

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("listen error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error", err)
		}

		go rpc.ServeConn(conn)
	}
}
