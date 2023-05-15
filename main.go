package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/guidogimeno/evm/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "the server address")
	flag.Parse()

	server := api.Server{ListenAddr: *listenAddr}
	fmt.Println("server running on port:", *listenAddr)

	log.Fatal(server.Start())
}
