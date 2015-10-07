package main

import (
	"echo"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type echoHandler struct{}

func (p *echoHandler) Echo(msg *echo.Message) (r string, err error) {
	fmt.Println("GoServer: " + msg.Text)
	return msg.Text, nil
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) {
	// Set processor
	handler := &echoHandler{}
	processor := echo.NewEchoProcessor(handler)

	// Set transport
	transport, _ := thrift.NewTServerSocket(addr)

	// Build and start server
	fmt.Println("GoServer started on ", addr)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	server.Serve()
}

/*
Even though we're not using a TThreadedServer like in server.py,
NewTSimpleServer spawns a goroutine to handle each incoming request,
https://git1-us-west.apache.org/repos/asf?p=thrift.git;a=blob;f=lib/go/thrift/simple_server.go
*/
