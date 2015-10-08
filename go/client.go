package main

import (
	"echo"
	"fmt"
	"net"
	"strconv"
	"sync"

	thrift "github.com/samuel/go-thrift/thrift"
)

func handleClient(addr string, num int) {
	// Open independent client connection
	conn, _ := net.Dial("tcp", addr)
	t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)
	client := thrift.NewClient(t, false)
	defer client.Close()
	ec := echo.EchoClient{Client: client}

	for i := num; i > 0; i-- {
		// Make thrift call and output result
		txt := strconv.Itoa(i)
		res, _ := ec.Echo(&echo.Message{Text: &txt})
		fmt.Println("GoClient: " + res)
	}
}

func runClient(addr string, num int) {
	// Number of concurrent clients to run
	clientCount := 10

	var wg sync.WaitGroup
	wg.Add(clientCount)

	// Spawn client connections
	for i := clientCount; i > 0; i-- {
		go func(num int) {
			handleClient(addr, num)
			wg.Done()
		}(num)
	}

	wg.Wait()
}
