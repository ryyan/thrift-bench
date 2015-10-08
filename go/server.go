package main

import (
	"echo"
	"net"
	"net/rpc"

	thrift "github.com/samuel/go-thrift/thrift"
)

type echoHandler struct{}

func (p *echoHandler) Echo(msg *echo.Message) (r string, err error) {
	return *msg.Text, nil
}

func runServer(port string) {
	// Set processor
	rpc.RegisterName("Thrift", &echo.EchoServer{Implementation: new(echoHandler)})
	ln, _ := net.Listen("tcp", port)

	for {
		conn, _ := ln.Accept()
		go func(c net.Conn) {
			t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(c, 0), thrift.BinaryProtocol)
			rpc.ServeCodec(thrift.NewServerCodec(t))
		}(conn)
	}
}
