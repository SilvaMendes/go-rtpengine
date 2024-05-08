package rtpengine

import (
	"fmt"
	"net"

	"github.com/google/uuid"
)

type Engine struct {
	con   net.Conn
	ip    net.IP
	port  int
	dns   *net.Resolver
	proto string
	ng    int
}

// Gera o cookie do comando
func (r *Engine) GetCookie() string {
	return uuid.NewString()
}

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

// Abrir conexão com p proxy rtpengine
func (r *Engine) Conn() (net.Conn, error) {
	engine := r.ip.String() + ":" + fmt.Sprint(r.port)
	conn, err := net.Dial(r.proto, engine)
	if err != nil {
		fmt.Println(err.Error(), r.proto, engine)
		return nil, err
	}
	r.con = conn
	return r.con, nil

}
