package rtpengine

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Client represents a network client that interacts with an underlying Engine.
// It encapsulates connection details, logging capabilities, and timeout configurations.
type Client struct {
	*Engine                // Embedded Engine instance providing core functionalities.
	url     string         // Base URL of the remote service to connect to.
	port    int            // Port number used for the connection.
	log     zerolog.Logger // Logger instance for structured logging and diagnostics.
	timeout time.Duration  // Maximum duration allowed for operations before timing out.
}

// ClientOption defines a function type that modifies a Client instance.
// It allows for flexible configuration of the Client during initialization.
// Each option function receives a pointer to the Client and returns an error if the configuration fails.
type ClientOption func(c *Client) error

// NewClient creates and initializes a new Client instance using the provided Engine and optional configuration functions.
// It sets default values for the URL, port, logger, and timeout, and applies any additional ClientOption functions.
// It also establishes a connection based on the protocol defined in the Engine.
//
// Parameters:
//   - rtpengine: Pointer to an Engine instance used by the Client.
//   - options: Variadic list of ClientOption functions for custom configuration.
//
// Returns:
//   - *Client: A pointer to the initialized Client instance.
//   - error: An error if any configuration or connection step fails.
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

	if c.Engine.proto == "udp" {
		if _, err := c.Engine.ConnUDP(); err != nil {
			c.log.Warn().Msg("Error connecting to RTP engine proxy: " + err.Error())
			return c, err
		}
	} else {
		if _, err := c.Engine.Conn(); err != nil {
			c.log.Warn().Msg("Error connecting to RTP engine proxy: " + err.Error())
			return c, err
		}
	}

	c.log = c.log.Level(zerolog.InfoLevel)
	return c, nil
}

// WithClientPort allows setting a custom default port for the Client.
// It returns a ClientOption function that updates both the Client's port
// and the associated Engine's port.
//
// Parameters:
//   - port: The port number to be set.
//
// Returns:
//   - ClientOption: A function that applies the port configuration to the Client.
func WithClientPort(port int) ClientOption {
	return func(s *Client) error {
		s.port = port
		s.Engine.port = port
		return nil
	}
}

// WithClientHostname sets the default hostname for the Client and resolves its IPv4 address.
// It returns a ClientOption function that updates the Client's IP field based on the resolved address.
//
// Parameters:
//   - hostname: The hostname to resolve (e.g., "localhost", "example.com").
//
// Returns:
//   - ClientOption: A function that applies the hostname resolution to the Client.
func WithClientHostname(hostname string) ClientOption {
	return func(s *Client) error {
		lookup, err := net.ResolveIPAddr("ip4", hostname)
		if err != nil {
			s.log.Warn().Msg("Error resolving hostname")
		}
		s.ip = lookup.IP
		return nil
	}
}

// WithClientDns sets the DNS resolver for the RTP engine service and resolves its IP address.
// It returns a ClientOption function that configures a custom DNS resolver using Google's public DNS (8.8.8.8),
// performs an IPv4 lookup for the specified domain, and updates the Client's URL with the resolved IP.
//
// Parameters:
//   - dns: The domain name of the RTP engine service to resolve.
//
// Returns:
//   - ClientOption: A function that applies the DNS resolution and updates the Client's URL.
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

// WithClientIP sets the default IP address for the Client.
// It returns a ClientOption function that parses the provided IP string
// and assigns it to the Client's IP field.
//
// Parameters:
//   - host: A string representing the IP address to be used.
//
// Returns:
//   - ClientOption: A function that applies the IP configuration to the Client.
func WithClientIP(host string) ClientOption {
	return func(s *Client) error {
		s.ip = net.ParseIP(host)
		return nil
	}
}

// WithClientProto sets the default protocol for the Client.
// It returns a ClientOption function that assigns the specified protocol string
// to the Client's proto field.
//
// Parameters:
//   - proto: A string representing the desired protocol (e.g., "udp", "tcp").
//
// Returns:
//   - ClientOption: A function that applies the protocol configuration to the Client.
func WithClientProto(proto string) ClientOption {
	return func(s *Client) error {
		s.proto = proto
		return nil
	}
}

// WithClientTimeout sets the connection timeout duration for the Client.
// It returns a ClientOption function that converts the provided integer value
// into a time.Duration and assigns it to the Client's timeout field.
//
// Parameters:
//   - t: Timeout value in milliseconds.
//
// Returns:
//   - ClientOption: A function that applies the timeout configuration to the Client.
func WithClientTimeout(t int) ClientOption {
	return func(s *Client) error {
		s.timeout = time.Duration(time.Duration(t).Milliseconds())
		return nil
	}
}

// SetLogLevel updates the logging level of the Client's logger.
// This method allows dynamic adjustment of the verbosity of log output
// based on the provided level.
//
// Parameters:
//   - level: An int8 value representing the desired logging level.
//     Valid levels are defined by zerolog.Level (e.g., DebugLevel, InfoLevel, WarnLevel, ErrorLevel).
func (s *Client) SetLogLevel(level int8) {
	s.log.Level(zerolog.Level(level))
}

// Close terminates the active connection held by the Client.
// If a UDP connection is active, it closes that connection.
// Otherwise, it closes the standard connection.
//
// Returns:
//   - error: Any error encountered while closing the connection.
func (s *Client) Close() error {
	if s.conUDP != nil {
		return s.conUDP.Close()
	} else {
		return s.con.Close()
	}
}

// NewComando sends a command to the RTP engine and retrieves the corresponding response.
// It generates a unique cookie, sends the command using ComandoNG, and then attempts to read the response using RespostaNG.
//
// Parameters:
//   - comando: A pointer to a RequestRtp struct containing the command to be sent.
//
// Returns:
//   - *ResponseRtp: A pointer to the response received from the RTP engine.
//     If an error occurs during command execution or response retrieval,
//     an empty ResponseRtp instance is returned.
func (c *Client) NewComando(comando *RequestRtp) *ResponseRtp {
	cookie := c.GetCookie()
	resposta := &ResponseRtp{}
	err := c.ComandoNG(cookie, comando)

	if err != nil {
		return resposta
	}

	resposta, err = c.RespostaNG(cookie)

	if err != nil {
		return resposta
	}

	return resposta
}

// ComandoNG sends a command to the RTP engine formatted in bencode.
// It encodes the command along with a unique cookie, logs the operation,
// and writes the message to the appropriate connection (UDP or TCP).
//
// Parameters:
//   - cookie: A unique identifier used to correlate the command and its response.
//   - comando: A pointer to a RequestRtp struct containing the command details.
//
// Returns:
//   - error: An error if encoding fails or if the message cannot be sent over the network.
func (c *Client) ComandoNG(cookie string, comando *RequestRtp) error {
	menssagem, err := EncodeComando(cookie, comando)
	if err != nil {
		return err
	}

	c.log.Debug().Msg("cookie: " + cookie + " Comando: " + comando.Command)

	if c.conUDP != nil {
		if _, err := c.conUDP.Write(menssagem); err != nil {
			return err
		}
	} else {
		if _, err := c.con.Write(menssagem); err != nil {
			return err
		}
	}

	return nil
}

// RespostaNG receives and decodes the response from the ngcp-rtpengine server.
// It reads raw data from the active connection (UDP or TCP), waits briefly to ensure the response is ready,
// and then decodes the response using the provided cookie.
//
// Parameters:
//   - cookie: A unique identifier used to match the response with the original command.
//
// Returns:
//   - *ResponseRtp: A pointer to the decoded response object.
//   - error: An error if reading from the connection fails.
func (c *Client) RespostaNG(cookie string) (*ResponseRtp, error) {
	respostaRaw := make([]byte, 65536)
	var err error
	resposta := &ResponseRtp{}

	if c.conUDP != nil {
		time.Sleep(1 * time.Second)
		_, err = c.conUDP.Read(respostaRaw)
	} else {
		time.Sleep(1 * time.Second)
		_, err = c.con.Read(respostaRaw)
	}

	if err != nil {
		return resposta, err
	}

	resposta = DecodeResposta(cookie, []byte(respostaRaw))
	return resposta, nil
}
