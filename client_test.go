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
	require.NotNil(t, rtp.port)
	fmt.Println("Func:", t.Name(), "Valor:", rtp.port)
	t.Skipped()

}

func TestClientRequestNewClientWithClientHostname(t *testing.T) {
	rtp, err := NewClient(
		&Engine{
			ip: net.ParseIP("10.0.0.0"),
		},
		WithClientHostname("L5NB-JGZXMF3"),
		WithClientProto("udp"))

	require.Nil(t, err)
	require.NotNil(t, rtp.ip)
	fmt.Println("Func:", t.Name(), "Valor:", rtp.ip, "PASS")
	t.Skipped()
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
	t.Skipped()
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

func TestClientRequestClientPing(t *testing.T) {
	t.Run("TestComandoPing", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2222), WithClientProto("udp"), WithClientDns("webrtcsrvgcp.callbox.com.br"))
		require.Nil(t, err)
		require.NotNil(t, client.Engine.con)
		r := &RequestRtp{
			Command: string(Ping),
		}
		response := client.NewComando(r)
		require.NotNil(t, response)

		fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
		c.con.Close()
	})
}

func TestClientRequestComando(t *testing.T) {
	sdp := `v=0
o=- 1545997027 1 IN IP4 198.51.100.1
s=tester
t=0 0
m=audio 2000 RTP/AVP 0
c=IN IP4 198.51.100.1
a=sendrecv`

	t.Run("TestComandoOffer", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2222), WithClientProto("udp"), WithClientDns("webrtcsrvgcp.callbox.com.br"))
		require.Nil(t, err)

		r := &RequestRtp{
			Command: string(Offer),
			ParamsOptString: &ParamsOptString{
				FromTag:           "asdasdasd494894",
				ToTag:             "asdasdad7879",
				CallId:            "5464asdas",
				TransportProtocol: string(RTP_AVP),
				Sdp:               sdp,
			},
			ParamsOptReplace: &ParamsOptReplace{
				Username:    "cbx-teste",
				SessionName: "cbx-teste",
			},
		}
		response := client.NewComando(r)
		require.NotNil(t, response)
		fmt.Println(response.Sdp)
		fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
	})

}
