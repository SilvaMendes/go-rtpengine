package gortpengine

import "net"

type RtpEngine struct {
	ip          net.IP
	port        int
	dnsResolver *net.Resolver
	ngPort      int
}

type RtpEngineOption func(s *RtpEngine) error

func (r *RtpEngine) GetIP() net.IP {
	return r.ip
}

func (r *RtpEngine) GetPort() int {
	return r.port
}
