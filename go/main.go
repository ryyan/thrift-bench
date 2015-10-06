package main

import (
	"flag"
	"runtime"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	// Use maximum number of CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse command line arguments
	num := flag.Int("num", 1, "Number of client requests to make")
	server := flag.Bool("server", false, "Run server if provided; Run client by default")
	flag.Parse()

	// Set transport. Should match what is in server.py
	addr := "localhost:9090"
	// Buffer size set to 4096 to match Python's DEFAULT_BUFFER,
	// https://git1-us-west.apache.org/repos/asf?p=thrift.git;a=blob;f=lib/py/src/transport/TTransport.py
	transportFactory := thrift.NewTBufferedTransportFactory(4096)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	if *server {
		runServer(transportFactory, protocolFactory, addr)
	} else {
		runClient(transportFactory, protocolFactory, addr, *num)
	}
}
