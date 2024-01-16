package api

import "net"

type RequestManagerListenerFactory interface {
	NewRequestManagerListener(logger LoggerService, dataService DataService) RequestManagerListener
}

type RequestManagerListener interface {
	TcpListener() *net.TCPListener
	UdpListener() net.PacketConn
	Start() error
	Stop()
}
