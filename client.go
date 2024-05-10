package rtpengine

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	*Engine
	url     string
	port    int
	log     zerolog.Logger
	timeout time.Duration
}

type ClientOption func(c *Client) error

func NewClient(rtpengine *Engine, options ...ClientOption) (*Client, error) {
	c := &Client{
		Engine:  rtpengine,
		url:     rtpengine.GetIP().String(),
		port:    rtpengine.GetPort(),
		log:     log.Logger.With().Str("Client", "RTPEngine").Logger(),
		timeout: 10 * time.Second,
	}

	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	if c.url != "" && c.url != "<nil>" {
		c.ip = net.ParseIP(c.url)
	}

	if _, err := c.Engine.Conn(); err != nil {
		c.log.Warn().Msg("Erro ao conectar com o proxy rtpengine " + err.Error())
	}

	return c, nil
}

// WithClientPort Permite definir a porta padrão do client
func WithClientPort(port int) ClientOption {
	return func(s *Client) error {
		s.port = port
		s.Engine.port = port
		return nil
	}
}

// WithClientHostname Permite definir o nome do host padrão do client resolve o endereço ipv4 da maquina local.
func WithClientHostname(hostname string) ClientOption {
	return func(s *Client) error {
		lookup, err := net.ResolveIPAddr("ip4", hostname)
		if err != nil {
			s.log.Warn().Msg("Erro resolver name")
		}
		s.ip = lookup.IP
		return nil
	}
}

// WithClientDns Permite definir o dns do serviço do rtpengine a função resolve o ip do serviço.
func WithClientDns(dns string) ClientOption {
	return func(s *Client) error {
		domain := &net.Resolver{
			PreferGo:     false,
			StrictErrors: false,
			Dial: func(ctx context.Context, network string, address string) (net.Conn, error) {
				return net.Dial("udp", "8.8.8.8:53")
			},
		}
		s.dns = domain
		url, _ := s.dns.LookupIP(context.TODO(), "ip4", dns)
		for _, resolver := range url {
			s.url = resolver.String()
		}
		return nil
	}
}

// WithClientPort Permite definir o IP padrão do client
func WithClientIP(host string) ClientOption {
	return func(s *Client) error {
		s.ip = net.ParseIP(host)
		return nil
	}
}

// WithClientPort Permite definir o protocolo padrão do client
func WithClientProto(proto string) ClientOption {
	return func(s *Client) error {
		s.proto = proto
		return nil
	}
}

// Fechar conexão aberta.
func (s *Client) Close() error {
	return s.con.Close()
}

func (c *Client) NewComando(comando *RequestRtp) *ResponseRtp {
	cookie := c.GetCookie()
	err := c.ComandoNG(cookie, comando)
	if err != nil {
		return nil
	}

	Resposta, err := c.RespostaNG(cookie)

	if err != nil {
		return nil
	}
	return Resposta
}

// Comando NG formatado em bencode para rtpengine
func (c *Client) ComandoNG(cookie string, comando *RequestRtp) error {
	menssagem, err := EncodeComando(cookie, comando)
	if err != nil {
		return err
	}

	c.log.Debug().Msg("cookie: " + cookie + " Comando: " + comando.Command)

	if _, err := c.con.Write(menssagem); err != nil {
		return err
	}
	return nil
}

// Resposta do servidor ngcp-rtpengine
func (c *Client) RespostaNG(cookie string) (*ResponseRtp, error) {
	c.con.SetReadDeadline(time.Now().Add(c.timeout))
	respostaRaw := make([]byte, 65536)

	_, err := c.con.Read(respostaRaw)
	if err != nil {
		return nil, err
	}

	resposta := DecodeResposta(cookie, respostaRaw)
	return resposta, nil
}
