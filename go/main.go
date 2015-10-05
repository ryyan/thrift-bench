package main

import (
	"flag"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	num := flag.Int("num", 1, "Number of client requests to make")
	server := flag.Bool("server", false, "Run server if provided; By default runs client")
	flag.Parse()

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()
	addr := "localhost:9090"

	if *server {
		runServer(transportFactory, protocolFactory, addr)
	} else {
		runClient(transportFactory, protocolFactory, addr, *num)
	}
}
