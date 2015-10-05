package main

import (
	"echo"
	"fmt"
	"strconv"
	"sync"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func handleClient(client *echo.EchoClient, txt string) {
	msg := &echo.Message{Text: txt}
	ret, err := client.Echo(msg)
	if txt == ret {
		fmt.Println("GoClient: ", ret)
	} else if err != nil {
		fmt.Println("GoClient: ERROR from server " + err.Error())
	} else {
		fmt.Println("GoClient: ERROR for " + txt)
	}
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, num int) error {
	var transport thrift.TTransport
	transport, err := thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}

	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(num)

	for num > 0 {
		go func(i int) {
			handleClient(echo.NewEchoClientFactory(transport, protocolFactory), strconv.Itoa(i))
			wg.Done()
		}(num)
		num--
	}

	wg.Wait()
	return nil
}
