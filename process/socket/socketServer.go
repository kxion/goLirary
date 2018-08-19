package socket

import (
	"log"
	"net"
)

type SocketServer struct{}

func NewSocketServer() *SocketServer {
	return new(SocketServer)
}

func (s *SocketServer) CreateTcpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8094")
	if err != nil {
		log.Printf("SocketServer CreateTcpServer create error:%s", err.Error())
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Printf("SocketServer CreateTcpServer listener error:%s", err.Error())
	}

}
