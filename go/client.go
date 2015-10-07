package main

import (
	"echo"
	"fmt"
	"sync"

	thrift "git.apache.org/thrift.git/lib/go/thrift"
	uuid "github.com/satori/go.uuid"
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

	for num > 0 {
		// Make thrift call and output result
		msg := &echo.Message{Text: uuid.NewV4().String()}
		ret, err := client.Echo(msg)

		if msg.Text == ret {
			fmt.Println("GoClient: " + ret)
		} else if err != nil {
			fmt.Println("GoClient: ERROR from server " + err.Error())
		} else {
			fmt.Println("GoClient: ERROR for " + msg.Text)
		}

		num--
	}
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, num int) {
	var wg sync.WaitGroup
	wg.Add(clientCount)

	// Spawn client connections
	for i := clientCount; i > 0; i-- {
		go func(num int) {
			handleClient(transportFactory, protocolFactory, addr, num)
			wg.Done()
		}(num)
	}

	wg.Wait()
}
