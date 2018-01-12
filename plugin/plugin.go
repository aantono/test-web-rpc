package main

import (
	"fmt"
	"net/rpc"
	"net"
	"github.com/aantono/test-web-rpc/api"
	"log"
)

func main() {
	fmt.Println("I am a plugin")
	startPlugin()
}

func startPlugin() {
	fmt.Println("Plugin start")

	// register the Plugin type with RPC.
	p := &api.Plugin{}
	err := rpc.Register(p)

	fmt.Println("Plugin: starting listener")
	// Start the listener.
	p.Listener, err = net.Listen("tcp", "127.0.0.1:55555")
	if err != nil {
		log.Fatal("Cannot listen: ", err)
	}
	// Start serving requests to the default server.
	fmt.Println("Plugin: accepting requests")
	rpc.Accept(p.Listener)
}