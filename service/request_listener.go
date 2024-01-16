package service

import (
	"IProtect/api"
	"net"
)

type RequestManagerListenerFactory struct {
	api.RequestManagerListener
}

type RequestManagerListener struct {
	logger      api.LoggerService
	tcpListener *net.TCPListener
	udpListener net.PacketConn
	dataService api.DataService
}

func (r RequestManagerListenerFactory) NewRequestManagerListener(logger api.LoggerService, dataService api.DataService) api.RequestManagerListener {
	return &RequestManagerListener{logger: logger, dataService: dataService}
}

func (r *RequestManagerListener) Start() error {
	var err error
	tcpListener, err := net.Listen("tcp", ":0")
	if err != nil {
		r.logger.Error("Unable to create tcp listener : " + err.Error())
		return err
	}
	r.tcpListener = tcpListener.(*net.TCPListener)

	r.handleTcpPacket()
	r.logger.Info("TCP listener are ready!")

	r.udpListener, err = net.ListenPacket("udp", "0.0.0.0:0")
	if err != nil {
		r.logger.Error("Unable to create udp listener : " + err.Error())
		return err
	}

	r.handleUdpPacket()
	r.logger.Info("UDP listener are ready!")
	return nil
}

func (r *RequestManagerListener) Stop() {
	err := r.tcpListener.Close()
	if err != nil {
		r.logger.Error("Unable to close listener : " + err.Error())
		return
	}

	err = r.udpListener.Close()
	if err != nil {
		r.logger.Error("Unable to close udp listener : " + err.Error())
		return
	}
}

func (r *RequestManagerListener) TcpListener() *net.TCPListener {
	return r.tcpListener
}

func (r *RequestManagerListener) UdpListener() net.PacketConn {
	return r.udpListener
}

func (r *RequestManagerListener) handleTcpPacket() {
	for {
		remoteAddress := r.tcpListener.Addr()
		if r.dataService.GetAddressState(remoteAddress.String()) {
			r.logger.Info("Address " + remoteAddress.String() + " try to connect but is blocked!")
			continue
		}

		conn, err := r.TcpListener().Accept()
		if err != nil {
			r.logger.Error("Unable to accept connection from address : " + remoteAddress.String())
			continue
		}
		err = conn.Close()
		if err != nil {
			r.logger.Error("Unable to close connection for remote address : " + remoteAddress.String())
			continue
		}
	}
}

func (r *RequestManagerListener) handleUdpPacket() {
	for {
		buffer := make([]byte, 1024)

		_, addr, err := r.udpListener.ReadFrom(buffer)
		if err != nil {
			r.logger.Error("Unable to read UDP packet :" + err.Error())
			continue
		}

		if r.dataService.GetAddressState(addr.String()) {
			r.logger.Info("Address " + addr.String() + " try to connect but is blocked!")
			_, err = r.UdpListener().WriteTo([]byte("Connection refused!"), addr)
			if err != nil {
				r.logger.Error("Unable to write refused access to remote udp request : " + err.Error())
				continue
			}
			continue
		}
	}
}
