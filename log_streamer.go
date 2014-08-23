package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"os"
	"bytes"
)

func check(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

type LogstashMsg struct {
	Timestamp string `json:"@timestamp"`
	Logtype	  string `json:"type"`
	Message string `json:"message"`
	Host	  string `json:"host"`
}

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	defer context.Close()
	defer socket.Close()

	//    following subscribes to all messages
	socket.SetSubscribe("")
	socket.Connect("tcp://localhost:2112")
	fmt.Println("Staring subscriber.....")

	for {
		logMessage := &LogstashMsg{}
		msg, _ := socket.Recv(0)
		// Strip the bytes from the messge or you get the follwing error
		// error: invalid character '\x00' after top-level value
		err := json.Unmarshal(bytes.Trim(msg, "\x00"), &logMessage)
		check(err)
		fmt.Println("#####", logMessage.Host, logMessage.Timestamp, "#####")
		fmt.Println(logMessage.Message)
	}
}
