package socket

import (
	"log"
	"net"
	"time"
)

type SocketClient struct{}

func NewSocketClient() *SocketClient {
	return new(SocketClient)
}

func (s *SocketClient) CreateTcpClient() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8094", 10*time.Second)
	if err != nil {
		log.Printf("SocketClient CreateTcpClient error:%s", err.Error())
	}
}
