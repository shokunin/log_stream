package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"bytes"
	"flag"
)

var hostname string
var port int

func init() {
	flag.StringVar(&hostname, "hostname", "localhost", "hostname to subscribe to")
	flag.IntVar(&port, "port", 2112, "port to try to connect to")
	flag.Parse()
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
	socket.Connect(fmt.Sprintf("tcp://%s:%d", hostname, port))
	fmt.Println("Staring subscriber.....")

	for {
		logMessage := &LogstashMsg{}
		msg, _ := socket.Recv(0)
		// Strip the bytes from the messge or you get the follwing error
		// error: invalid character '\x00' after top-level value
		err := json.Unmarshal(bytes.Trim(msg, "\x00"), &logMessage)
		if err == nil {
			fmt.Println("#####", logMessage.Host, logMessage.Timestamp, "#####")
			fmt.Println(logMessage.Message)
		} else {
			fmt.Println("##### ERR:", err, "#####")
			fmt.Println(string(msg))
		}
	}
}
