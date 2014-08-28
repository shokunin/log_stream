package main

import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:2115")
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: %s\n", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("ListenTCP failed: %s\n", err.Error())
	}

	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUB)
	defer context.Close()
	defer socket.Close()
	socket.Bind("tcp://*:2112")
	// run a socket if necessary
	// socket.Bind("ipc://logs.ipc")

	for {
		conn, err := listener.Accept()
		message := make([]byte, 4096)
		n, err := conn.Read(message)
		// fmt.Println(string(message))
		socket.Send([]byte(message), 0)
		if err != nil || n == 0 {
			fmt.Println("Error is: %s, bytes are: %d", err, n)
			continue
		}
	}
}
