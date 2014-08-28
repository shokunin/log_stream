all:
	go build udp2zmq.go
	go build tcp2zmq.go
	go build log_streamer.go

log_streamer:
	go build log_streamer.go

tcp:
	go build tcp2zmq.go

udp:
	go build udp2zmq.go
