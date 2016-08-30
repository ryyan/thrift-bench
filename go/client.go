package main

import (
	"echo"
	"fmt"
	"net"
	"strconv"
	"sync"
	"sync/atomic"

	thrift "github.com/samuel/go-thrift/thrift"
	uuid "github.com/satori/go.uuid"
)

var count uint64

func handleClient(addr string, num int) {
	// Open independent client connection
	conn, _ := net.Dial("tcp", addr)
	t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)
	client := thrift.NewClient(t, false)
	defer client.Close()
	ec := echo.EchoClient{Client: client}

	// UUID
	uid := uuid.NewV4().String()

	for i := 0; i < num; i++ {
		// Make thrift call and increment atomic count
		txt := uid + strconv.Itoa(i)
		ret, _ := ec.Echo(&echo.Message{Text: &txt})
		if txt == ret {
			atomic.AddUint64(&count, 1)
		}
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

	if int(count) != (clientCount * num) {
		fmt.Println("ERROR: Actual and expected completed requests mismatch")
		fmt.Println("Expected: " + strconv.Itoa(clientCount*num))
		fmt.Println("Actual: " + strconv.FormatUint(count, 10))
	}
}
