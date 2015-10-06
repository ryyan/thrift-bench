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

func handleClient(c *client, num int) {
	// Open independent client connection
	client := c.open()
	defer c.close()

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

	for num > 0 {

		go func(i int) {
			// Pass a new client for each request
			client := client{transportFactory, protocolFactory, addr, nil}
			handleClient(&client, i)
			wg.Done()
		}(num)

		num--
	}

	wg.Wait()
}

// Helper class used to build independent client connections
type client struct {
	transportFactory thrift.TTransportFactory
	protocolFactory  thrift.TProtocolFactory
	addr             string
	client           *echo.EchoClient
}

// Build and open a new client connection
func (c *client) open() *echo.EchoClient {
	var transport thrift.TTransport
	transport, _ = thrift.NewTSocket(c.addr)
	transport = c.transportFactory.GetTransport(transport)

	c.client = echo.NewEchoClientFactory(transport, c.protocolFactory)
	c.client.Transport.Open()
	return c.client
}

// Close connection
func (c *client) close() {
	c.client.Transport.Close()
}
