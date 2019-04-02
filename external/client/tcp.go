package client

import (
	"net"
)

// TCPClient represents the TCP client.
type TCPClient struct {
	net.Conn
}

// NewTCPClient returns "TCPClient".
func NewTCPClient(addr string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpClt := &TCPClient{Conn: conn}
	return tcpClt, nil
}

// Write write a message to a device.
func (c *TCPClient) Write(msg string) error {
	byteMsg := []byte(msg)
	if _, err := c.Conn.Write(byteMsg); err != nil {
		return err
	}
	return nil
}

// Read read a buffer from a device.
func (c *TCPClient) Read(bufSize int) ([]byte, error) {
	buf := make([]byte, bufSize)
	bufLen, err := c.Conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:bufLen], nil
}

// Query query to a device.
func (c *TCPClient) Query(msg string, bufSize int) ([]byte, error) {
	if err := c.Write(msg); err != nil {
		return nil, err
	}
	buf, err := c.Read(bufSize)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
