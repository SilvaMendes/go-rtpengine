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
			Command:              string(Offer),
			ParamsOptString:      &ParamsOptString{FromTag: "asdasdasd494894AAAA", ToTag: "asdasdad7879000", CallId: "5464asdas00000000", TransportProtocol: string(RTP_AVP), Sdp: sdp},
			ParamsOptStringArray: &ParamsOptStringArray{Replace: []ParamReplace{Username, SessionName}},
		}
		response := client.NewComando(r)
		require.NotNil(t, response)
		fmt.Println(response.Sdp)
		fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
	})
	time.Sleep(4 * time.Second)
	t.Run("TestComandoDelete", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2222), WithClientProto("udp"), WithClientDns("webrtcsrvgcp.callbox.com.br"))
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
		require.NotNil(t, response.Sdp)
		fmt.Println(response.Sdp)
		fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
	})

}

func TestClientRequestOffer(t *testing.T) {
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

		r := &ParamsOptString{FromTag: "asdasdasd494894", ToTag: "asdasdad7879", CallId: "5464asdas", TransportProtocol: string(RTP_AVP), Sdp: sdp}
		flags := make([]string, 0)
		flags = append(flags, string(TrustAddress))

		repl := make([]ParamReplace, 0)
		repl = append(repl, SessionConnection, Origin)

		rtcpmux := make([]ParamRTCPMux, 0)
		rtcpmux = append(rtcpmux, RTCP_Demux, RTCP_Accept)
		opt := &RequestRtp{}

		transcoderList := make([]string, 0)
		transcoderList = append(transcoderList, string(CODEC_G722))

		request, err := SDPOffering(r, opt.SetFlags(flags), opt.SetTransportProtocol(RTP_AVP), opt.SetReplace(repl), opt.SetRtcpMux(rtcpmux), opt.SetCodecEncoder(transcoderList))
		require.Nil(t, err)
		response := client.NewComando(request)
		require.NotNil(t, response.Sdp)
		fmt.Println(response.Sdp)
		fmt.Println("Func:", t.Name(), "Comando:"+request.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
	})

	time.Sleep(4 * time.Second)
	t.Run("TestComandoDelete", func(t *testing.T) {
		c := &Engine{}
		client, err := NewClient(c, WithClientPort(2222), WithClientProto("udp"), WithClientDns("webrtcsrvgcp.callbox.com.br"))
		require.Nil(t, err)

		r := &RequestRtp{
			Command: string(Delete),
			ParamsOptString: &ParamsOptString{
				FromTag: "asdasdasd494894",
				ToTag:   "asdasdad7879",
				CallId:  "5464asdas",
			},
		}

		response := client.NewComando(r)
		require.NotNil(t, response)
		fmt.Println("Func:", t.Name(), "Comando:"+r.Command, "Resposta:"+response.Result, "Motivo:", response.ErrorReason, client.con.RemoteAddr().String(), "PASS")
	})
}
