package main

import (
	"flag"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	// Parse command line arguments
	num := flag.Int("num", 1, "Number of client requests to make")
	server := flag.Bool("server", false, "Run server if provided; Run client by default")
	flag.Parse()

	// Set transport
	addr := "localhost:9090"
	transportFactory := thrift.NewTBufferedTransportFactory(4096)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	if *server {
		runServer(transportFactory, protocolFactory, addr)
	} else {
		runClient(transportFactory, protocolFactory, addr, *num)
	}
}

/*
Buffer size set to 4096 to match Python's DEFAULT_BUFFER,
https://git1-us-west.apache.org/repos/asf?p=thrift.git;a=blob;f=lib/py/src/transport/TTransport.py
*/
