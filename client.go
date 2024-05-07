package gortpengine

import (
	"context"
	"net"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	*RtpEngine
	url  string
	port int
	log  zerolog.Logger
}

type ClientOption func(c *Client) error

func NewClient(rtpengine *RtpEngine, options ...ClientOption) (*Client, error) {
	c := &Client{
		RtpEngine: rtpengine,
		url:       rtpengine.GetIP().String(),
		port:      rtpengine.GetPort(),
		log:       log.Logger.With().Str("New", "Client").Logger(),
	}

	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithClientPort permite definir a porta padrão do client
func WithClientPort(port int) ClientOption {
	return func(s *Client) error {
		s.port = port
		s.RtpEngine.port = port
		return nil
	}
}

// WithClientHostname permite definir o nome do host padrão do client resolve o endereço ipv4 da maquina local.
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

// WithClientDns permite definir o dns do serviço do rtpengine a função resolve o ip do serviço.
func WithClientDns(dns string) ClientOption {
	return func(s *Client) error {
		domain := &net.Resolver{
			PreferGo:     false,
			StrictErrors: false,
			Dial: func(ctx context.Context, network string, address string) (net.Conn, error) {
				return net.Dial("udp", "8.8.8.8:53")
			},
		}
		s.dnsResolver = domain
		url, _ := s.dnsResolver.LookupIP(context.TODO(), "ip4", dns)
		for _, resolver := range url {
			s.url = resolver.String()
		}
		return nil
	}
}
