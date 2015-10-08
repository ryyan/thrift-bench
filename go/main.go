package main

import (
	"flag"
)

func main() {
	// Parse command line arguments
	num := flag.Int("num", 1, "Number of requests each client will make")
	server := flag.Bool("server", false, "Run server if provided; Run client by default")
	flag.Parse()
	port := ":9999"

	if *server {
		runServer(port)
	} else {
		addr := "127.0.0.1" + port
		runClient(addr, *num)
	}
}
