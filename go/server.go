package main

import (
	"echo"
	"errors"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (p *EchoHandler) Echo(msg *echo.Message) (r string, err error) {
	if msg.Text == "" {
		err = errors.New("No text supplied")
		return
	}
	fmt.Println("GoServer: " + msg.Text)
	return msg.Text, nil
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	fmt.Printf("%T\n", transport)
	processor := echo.NewEchoProcessor(NewEchoHandler())
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
