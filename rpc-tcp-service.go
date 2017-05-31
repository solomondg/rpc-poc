package main

import (
	"github.com/solomondg/rpc-poc/remote"
	"net"
	"net/rpc"
)

func main() {
	pi := new(remote.Pi)

	rpc.Register(pi)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}

	rpc.Accept(listener)
}
