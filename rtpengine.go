package rtpengine

import (
	"bytes"
	"fmt"
	"net"

	bencode "github.com/anacrolix/torrent/bencode"
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

// Trasformar o comando em bencode
func EncodeComando(cookie string, command *RequestRtp) ([]byte, error) {
	data, err := bencode.Marshal(command)
	if err != nil {
		return nil, err
	}

	bind := []byte(cookie + " ")
	return append(bind, data...), nil
}

func DecodeResposta(cookie string, resposta []byte) *ResponseRtp {
	resp := &ResponseRtp{}
	cookieIndex := bytes.IndexAny(resposta, " ")
	if cookieIndex != len(cookie) {
		resp.Result = "error"
		resp.ErrorReason = "Erro ao analisar a mensagem"
		return resp
	}

	cookieResponse := string(resposta[:cookieIndex])
	if cookieResponse != cookie {
		resp.Result = "error"
		resp.ErrorReason = "O cookie não corresponde"
		return resp
	}

	encodedData := string(resposta[cookieIndex+1:])
	err := bencode.Unmarshal([]byte(encodedData), resp)

	if err != nil {
		return resp
	}

	return resp
}
