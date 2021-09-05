package main

import (
	"log"
	"net/rpc"
	"os"
)

type Args struct {}

func main() {
    // DialHTTP connects to an HTTP RPC server at the specified network
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var reply int64
	args := Args{}

	err = client.Call("TimeServer.Current", args, &reply)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("%d", reply)
}
