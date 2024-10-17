package rtpengine

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientRequestNewClientWithClientPort(t *testing.T) {
	rtp, err := NewClient(
		&Engine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientPort(2222),
		WithClientProto("udp"))
	require.Nil(t, err)
	require.NotNil(t, &rtp.port)
	fmt.Println("Func:", t.Name(), "Valor:", rtp.port)

}

func TestClientRequestNewClientWithClientHostname(t *testing.T) {
	rtp, err := NewClient(
		&Engine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientHostname("L5NB-JGZXMF3"),
		//WithClientHostname("DESKTOP-QJ365M6"),
		WithClientProto("udp"))

	require.Nil(t, err)
	require.NotNil(t, rtp.ip)
	fmt.Println("Func:", t.Name(), "Valor:", rtp.ip, "PASS")
}

func TestClientRequestNewClienWithClientDns(t *testing.T) {
	rtp, err := NewClient(
		&Engine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientDns("webrtcsrvgcp.callbox.com.br"),
		WithClientProto("udp"))

	require.Nil(t, err)
	require.NotNil(t, rtp.url)
	fmt.Println("Func:", t.Name(), "Valor:", rtp.url, "PASS")
}

func TestClientRequestClientOption(t *testing.T) {
	t.Run("TestClientDNS", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2222), WithClientProto("udp"), WithClientDns("webrtcsrvgcp.callbox.com.br"))
		require.Nil(t, err)
		require.NotNil(t, client.Engine.con)
		fmt.Println("Func:", t.Name(), "Valor:", client.con.RemoteAddr().String(), "PASS")
		c.con.Close()
	})

	t.Run("TestClientIP", func(t *testing.T) {
		b := &Engine{}
		clt, err := NewClient(b, WithClientPort(2222), WithClientProto("udp"), WithClientIP("35.198.31.42"))
		require.Nil(t, err)
		require.NotNil(t, clt.Engine.con)
		fmt.Println("Func:", t.Name(), "Valor:", clt.con.RemoteAddr().String(), "PASS")
		b.con.Close()
	})
}
