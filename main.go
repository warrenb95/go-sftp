package main

import (
	"log"
	"net"

	"github.com/pkg/sftp"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:2022")
	if err != nil {
		log.Fatalf("Failed to listen on port 2022, err: %v", err)
	}

	netConn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connections on listener, err: %v", err)
	}

	s, err := sftp.NewServer(netConn)
	if err != nil {
		log.Fatalf("Failed to create new sftp server, err: %v", err)
	}

	err = s.Serve()
	if err != nil {
		log.Fatalf("Failed to server sftp server, err: %v", err)
	}
}
