package rtpengine

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClientRequestClientPing(t *testing.T) {
	t.Run("TestComandoPing", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2221), WithClientProto("tcp"), WithClientDns("rtp01"))
		require.Nil(t, err)
		if client.Engine.conUDP != nil {
			require.NotNil(t, client.Engine.conUDP)
		} else {
			require.NotNil(t, client.Engine.con)

		}
		r := &RequestRtp{
			Command: string(Ping),
		}

		response := client.NewComando(r)
		client.Close()
		require.NotNil(t, response.Result)
		if client.conUDP != nil {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.conUDP.RemoteAddr().String(), "PASS")
		} else {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
		}

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
		client, err := NewClient(c, WithClientPort(2221), WithClientProto("tcp"), WithClientDns("rtp01"))
		require.Nil(t, err)

		r := &RequestRtp{
			Command:              string(Offer),
			ParamsOptString:      &ParamsOptString{FromTag: "asdasdasd494894AAAA", ToTag: "asdasdad7879000", CallId: "5464asdas00000000", TransportProtocol: RTP_AVP, Sdp: sdp},
			ParamsOptStringArray: &ParamsOptStringArray{Replace: []ParamReplace{Username, SessionName}},
		}
		response := client.NewComando(r)
		client.Close()

		require.NotNil(t, response.Sdp)
		fmt.Println(response.Sdp)
		if client.conUDP != nil {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.conUDP.RemoteAddr().String(), "PASS")
		} else {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
		}
	})
	time.Sleep(2 * time.Second)

	t.Run("Query", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2221), WithClientProto("tcp"), WithClientDns("rtp01"))
		require.Nil(t, err)

		r := &RequestRtp{
			Command:         string(Query),
			ParamsOptString: &ParamsOptString{CallId: "5464asdas00000000"},
		}

		response := client.NewComando(r)
		client.Close()

		require.NotNil(t, response.Result)
		if tagsMap, ok := response.Tags.(map[string]interface{}); ok {
			for k, value := range tagsMap {
				fmt.Println(k, " => ", value)
			}
		}

		if client.conUDP != nil {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.conUDP.RemoteAddr().String(), "PASS")
		} else {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
		}
	})

	time.Sleep(2 * time.Second)
	t.Run("TestComandoDelete", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2221), WithClientProto("tcp"), WithClientDns("rtp01"))
		require.Nil(t, err)

		r := &RequestRtp{
			Command: string(Delete),
			ParamsOptString: &ParamsOptString{
				FromTag: "asdasdasd494894AAAA",
				ToTag:   "asdasdad7879000",
				CallId:  "5464asdas00000000",
			},
		}

		response := client.NewComando(r)
		client.Close()

		require.NotNil(t, response.Result)
		if client.conUDP != nil {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.conUDP.RemoteAddr().String(), "PASS")
		} else {
			fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
		}
	})
}
