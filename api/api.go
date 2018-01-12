package api

import (
	"net"
	"fmt"
	"os"
)

type Plugin struct {
	Listener net.Listener
}

func (p Plugin) Reverse(arg string, ret *string) error {
	fmt.Println("Plugin: reverse")
	l := len(arg)
	r := make([]byte, l)
	for i := 0; i < l; i++ {
		r[i] = arg[l-1-i]
	}
	*ret = string(r)
	return nil
}

func (p Plugin) Exit(arg int, ret *int) error {
	fmt.Println("Plugin: done.")
	os.Exit(0) // using os.Exit here is suitable for demo code only.
	return nil
}