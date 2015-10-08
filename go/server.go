package main

import (
	"echo"
	"fmt"
	"net"
	"net/rpc"

	thrift "github.com/samuel/go-thrift/thrift"
)

type echoHandler struct{}

func (p *echoHandler) Echo(msg *echo.Message) (r string, err error) {
	fmt.Println("GoServer: " + *msg.Text)
	return *msg.Text, nil
}

func runServer(port string) {
	// Set processor
	rpc.RegisterName("Thrift", &echo.EchoServer{Implementation: &echoHandler{}})
	ln, _ := net.Listen("tcp", port)

	for {
		conn, _ := ln.Accept()
		t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)
		go rpc.ServeCodec(thrift.NewServerCodec(t))
	}
}
