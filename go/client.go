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

// Opens independent client connection and talks to thrift server
func handleClient(c *client, num int) {
	client := c.open()
	defer c.close()

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

// Spawns client goroutines
func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, num int) {
	var wg sync.WaitGroup
	wg.Add(num)

	for num > 0 {
		go func(i int) {
			client := client{transportFactory, protocolFactory, addr, nil}
			handleClient(&client, i)
			wg.Done()
		}(num)
		num--
	}
	wg.Wait()
}

// Helper object used to build independent client connections
type client struct {
	transportFactory thrift.TTransportFactory
	protocolFactory  thrift.TProtocolFactory
	addr             string
	client           *echo.EchoClient
}

func (c *client) open() *echo.EchoClient {
	var transport thrift.TTransport
	transport, _ = thrift.NewTSocket(c.addr)
	transport = c.transportFactory.GetTransport(transport)

	c.client = echo.NewEchoClientFactory(transport, c.protocolFactory)
	c.client.Transport.Open()
	return c.client
}

func (c *client) close() {
	c.client.Transport.Close()
}
