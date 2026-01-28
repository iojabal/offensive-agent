package main

import (
	"nombredetuapp/Documents/Proyecto/src/Dispatcher"
	"nombredetuapp/Documents/Proyecto/src/transport"
)

func main() {
	var conn transport.TCPTransport
	var err error

	conn = *transport.NewTCPTransport("192.168.148.151:443")
	err = conn.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	Dispatcher.Run(conn)
}
