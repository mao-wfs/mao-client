package client

import (
	"net"
)

// TCPClient is the structure of a TCP client.
type TCPClient struct {
	Conn net.Conn
}

// NewTCPClient returns the new TCP client.
func NewTCPClient(addr string) TCPClient {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	tcpClient := TCPClient{
		Conn: conn,
	}
	return tcpClient
}
