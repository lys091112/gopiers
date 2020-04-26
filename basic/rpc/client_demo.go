package rpc

import (
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

// 检测 client也满足 FirstServiceInterface 接口
// TODO 仅用于检测，也可以没有
var _ FristInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(FirstServiceName+".Hello", request, reply)
}

func StartClient() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}
