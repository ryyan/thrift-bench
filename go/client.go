package main

import (
	"echo"
	"fmt"
	"strconv"
	"sync"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// Number of concurrent clients to run
const clientCount = 10

func handleClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, num int) {
	// Open independent client connection
	var transport thrift.TTransport
	transport, _ = thrift.NewTSocket(addr)
	transport = transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(transport, protocolFactory)
	client.Transport.Open()
	defer client.Transport.Close()

	// Make thrift call and output result
	msg := &echo.Message{Text: strconv.Itoa(num)}
	ret, err := client.Echo(msg)
	if msg.Text == ret {
		fmt.Println("GoClient: " + ret)
	} else if err != nil {
		fmt.Println("GoClient: ERROR from server " + err.Error())
	} else {
		fmt.Println("GoClient: ERROR for " + msg.Text)
	}
}

// Spawns client connections
func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, num int) {
	var wg sync.WaitGroup
	wg.Add(num)

	for i := num; i > 0; i-- {
		go func(i int) {
			handleClient(transportFactory, protocolFactory, addr, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
