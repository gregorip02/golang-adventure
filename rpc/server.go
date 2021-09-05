package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type Args struct {}
type TimeServer int64

func (t *TimeServer) Current(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	current := new(TimeServer)
	rpc.Register(current)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Listening on port 3000")
	http.Serve(listener, nil)
}
