package main

import (
	"fmt"
	"net/rpc"

	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("I am a main app...")
	startApp()
}

func startApp() {
	fmt.Println("App start")
	// Get the plugin.
	p := exec.Command("../plugin/plugin")
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr

	// Start the plugin process.
	err := p.Start()
	if err != nil {
		log.Fatal("Cannot start ", p.Path, ": ", err)
	}
	// Ensure the plugin process is up and running before attempting to connect.
	// `net/rpc` has no `DialTimeout` method (unlike `net`).
	time.Sleep(1 * time.Second)

	// Create the RPC client.
	fmt.Println("App: registering RPC client")

	client, err := rpc.Dial("tcp", "127.0.0.1:55555")
	if err != nil {
		log.Fatal("Cannot create RPC client: ", err)
	}

	// Call the Revert method.
	fmt.Println("App: calling Reverse")

	for i := 0; i < 500; i++ {
		t := time.Now()
		var reverse string
		err = client.Call("Plugin.Reverse", "Live on time, emit no evil", &reverse)
		if err != nil {
			log.Fatal("Error calling Reverse: ", err)
		}
		d := time.Now().Sub(t)
		fmt.Println("App: revert result:", reverse)
		fmt.Println("App: took " + d.String())
	}

	// Stop the plugin and terminate the app.
	fmt.Println("App: stopping the plugin")
	var n int
	client.Call("Plugin.Exit", 0, &n)
	p.Wait()
	fmt.Println("App: done.")
}