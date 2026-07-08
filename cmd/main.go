package main

import (
	"dakt/cmd/forward"
	"flag"
)

func main() {
	port := flag.String("port", "8080", "Proxy port")
	flag.Parse()

	forward.Forward(*port)
}
