package rtpengine

import (
	"net"
)

type Engine struct {
	ip    net.IP
	port  int
	dns   *net.Resolver
	proto string
	ng    int
}

type RtpEngineOption func(s *Engine) error

// Atribuir o ip padrão para conexão
func (r *Engine) GetIP() net.IP {
	return r.ip
}

// Atribuir a porta padrão para conexão
func (r *Engine) GetPort() int {
	return r.port
}

// Atribuir o protocolo padrão para conexão
func (r *Engine) GetProto() string {
	return r.proto
}

// Atribuir a porta padrão NG porta de controler
func (r *Engine) GetNG() int {
	return r.ng
}
