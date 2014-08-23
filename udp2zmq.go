package main

import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"log"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:2112")
	if err != nil {
		log.Fatalf("ResolveUDPAddr failed: %s\n", err)
	}
	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("ListenUDP failed: %s\n", err.Error())
	}

	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUB)
	defer context.Close()
	defer socket.Close()
	socket.Bind("tcp://*:2112")
	// run a socket if necessary
	// socket.Bind("ipc://logs.ipc")

	for {
		message := make([]byte, 1024)
		n, _, err := listener.ReadFromUDP(message)
		fmt.Println(string(message))
		socket.Send([]byte(message), 0)
		if err != nil || n == 0 {
			log.Printf("Error is: %s, bytes are: %d", err, n)
			continue
		}
	}
}
