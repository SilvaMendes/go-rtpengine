package gortpengine

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientRequestNewClientWithClientPort(t *testing.T) {
	rtp, err := NewClient(
		&RtpEngine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientPort(2222))
	require.Nil(t, err)
	require.NotNil(t, rtp.port)
	fmt.Println("Func:", t.Name(), "Value:", rtp.port)

}

func TestClientRequestNewClientWithClientHostname(t *testing.T) {
	rtp, err := NewClient(
		&RtpEngine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientHostname("L5NB-JGZXMF3"))

	require.Nil(t, err)
	require.NotNil(t, rtp.ip)
	fmt.Println("Func:", t.Name(), "Value:", rtp.ip)
}

func TestClientRequestNewClienWithClientDns(t *testing.T) {
	rtp, err := NewClient(
		&RtpEngine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientDns("webrtcsrvgcp.callbox.com.br"))

	require.Nil(t, err)
	require.NotNil(t, rtp.url)
	fmt.Println("Func:", t.Name(), "Value:", rtp.url)
}

func TestClientRequestClientOption(t *testing.T) {
	var opt ClientOption = func(c *Client) error {
		WithClientPort(2222)
		WithClientDns("webrtcsrvgcp.callbox.com.br")
		return nil
	}
	c := &Client{}
	rtp := &RtpEngine{
		ip:          c.GetIP(),
		port:        c.GetPort(),
		dnsResolver: &net.Resolver{},
		ngPort:      0,
	}
	client, err := NewClient(rtp, opt)
	require.Nil(t, err)
	fmt.Println(client)
}
