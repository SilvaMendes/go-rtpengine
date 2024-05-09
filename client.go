package rtpengine

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stefanovazzocell/bencode"
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

func (c *Client) Comando(comando string) (map[string]interface{}, error) {
	cookie := c.GetCookie()
	err := c.enviaComando(cookie, comando)
	if err != nil {
		return nil, err
	}

	decodedMsg, err := c.retornaResposta(cookie)
	if err != nil {
		return nil, err
	}
	return decodedMsg, nil
}

func (c *Client) enviaComando(cookie string, command string) error {
	message, err := encodeCommand(cookie, command)
	if err != nil {
		return err
	}
	c.log.Debug().Msg("cookie: " + cookie + " Comando: " + command)
	if _, err := c.con.Write(message); err != nil {
		return err
	}
	return nil
}

func (c *Client) retornaResposta(cookie string) (map[string]interface{}, error) {
	c.con.SetReadDeadline(time.Now().Add(c.timeout))
	response := make([]byte, 65536)
	_, err := c.con.Read(response)
	if err != nil {
		return nil, err
	}

	decodedMsg, err := decodeResponse(cookie, response)

	if err != nil {
		return nil, err
	}
	c.log.Debug().Msg("Cookie: " + cookie + " Resposta: " + fmt.Sprint(decodedMsg["result"]))

	return c.validateResponse(decodedMsg)
}

func (c *Client) validateResponse(decodedMsg map[string]interface{}) (map[string]interface{}, error) {
	result, ok := decodedMsg["result"]
	if !ok {
		return nil, errors.New("Sem retorno do Result")
	}
	if result == "ok" || result == "pong" {
		if err, ok := decodedMsg["warning"]; ok {
			fmt.Println(err)
			c.log.Warn().Msg("Error")
		}
		return decodedMsg, nil
	}
	if result == "error" {
		if reason, ok := decodedMsg["error-reason"]; ok {
			return nil, errors.New(reason.(string))
		}
	}
	return nil, errors.New("Error Desconhecido")
}

func decodeResponse(cookie string, response []byte) (map[string]interface{}, error) {
	cookieIndex := bytes.IndexAny(response, " ")
	if cookieIndex != len(cookie) {
		return nil, errors.New(" Erro ao analisar a mensagem")
	}

	cookieResponse := string(response[:cookieIndex])
	if cookieResponse != cookie {
		return nil, errors.New("O cookie não corresponde")
	}

	encodedData := string(response[cookieIndex+1:])
	decodedData, err := bencode.NewParserFromString(encodedData).AsDict()
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

func encodeCommand(cookie string, command string) ([]byte, error) {
	var dict interface{} = map[string]interface{}{
		"command": command,
	}
	bdata, err := bencode.NewEncoderFromInterface(dict)
	if err != nil {
		return nil, err
	}
	bid := []byte(cookie + " ")
	return append(bid, bdata.Bytes()...), nil
}
