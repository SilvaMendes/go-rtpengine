package gortpengine

import (
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
		url:       "",
		port:      0,
		log:       log.Logger.With().Str("New", "Client").Logger(),
	}

	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
