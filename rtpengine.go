package rtpengine

import (
	"bytes"
	"fmt"
	"net"
	"time"

	bencode "github.com/anacrolix/torrent/bencode"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	ben "github.com/stefanovazzocell/bencode"
)

type Engine struct {
	con    net.Conn
	conUDP *net.UDPConn
	ip     net.IP
	port   int
	dns    *net.Resolver
	proto  string
	ng     int
	*ResponseRtp
}

// Estrutura da requisicão do comando
type RequestRtp struct {
	Command string `json:"command" bencode:"command"`
	*ParamsOptString
	*ParamsOptInt
	*ParamsOptStringArray
}

// Estrutura da resposta do comando
type ResponseRtp struct {
	Result          string      `json:"result" bencode:"result"`
	Sdp             string      `json:"sdp,omitempty" bencode:"sdp,omitempty"`
	ErrorReason     string      `json:"error-reason,omitempty" bencode:"error-reason,omitempty"`
	Warning         string      `json:"warning,omitempty" bencode:"warning,omitempty"`
	Created         int         `json:"created,omitempty" bencode:"created,omitempty"`
	CreatedUs       int         `json:"created_us,omitempty" bencode:"created_us,omitempty"`
	LastSignal      int         `json:"last signal,omitempty" bencode:"last signal,omitempty"`
	LastRedisUpdate int         `json:"last redis update,omitempty" bencode:"last redis update,omitempty"`
	SSRC            interface{} `json:"SSRC,omitempty" bencode:"SSRC,omitempty"`
	Tags            interface{} `json:"tags,omitempty" bencode:"tags,omitempty"`
	Totals          TotalRTP    `json:"totals,omitempty" bencode:"totals,omitempty"`
}

type TotalRTP struct {
	Rtp  ValuesRTP `json:"RTP,omitempty" bencode:"RTP,omitempty"`
	Rtcp ValuesRTP `json:"RCTP,omitempty" bencode:"RTP,omitempty"`
}
type ValuesRTP struct {
	Packets int `json:"packets,omitempty" bencode:"packets,omitempty"`
	Bytes   int `json:"bytes,omitempty" bencode:"bytes,omitempty"`
	Errors  int `json:"errors,omitempty" bencode:"errors,omitempty"`
}

// Parametros de comportamento
type ParamsOptString struct {
	FromTag                string                 `json:"from-tag,omitempty" bencode:"from-tag,omitempty"`
	ToTag                  string                 `json:"to-tag,omitempty" bencode:"to-tag,omitempty"`
	CallId                 string                 `json:"call-id,omitempty" bencode:"call-id,omitempty"`
	TransportProtocol      TransportProtocol      `json:"transport-protocol,omitempty" bencode:"transport-protocol,omitempty"`
	MediaAddress           string                 `json:"media-address,omitempty" bencode:"media-address,omitempty"`
	ICE                    ICE                    `json:"ICE,omitempty" bencode:"ICE,omitempty"`
	AddressFamily          AddressFamily          `json:"address-family,omitempty" bencode:"address-family,omitempty"`
	DTLS                   DTLS                   `json:"DTLS,omitempty" bencode:"DTLS,omitempty"`
	ViaBranch              string                 `json:"via-branch,omitempty" bencode:"via-branch,omitempty"`
	XmlrpcCallback         string                 `json:"xmlrpc-callback,omitempty" bencode:"xmlrpc-callback,omitempty"`
	Metadata               string                 `json:"metadata,omitempty" bencode:"metadata,omitempty"`
	File                   string                 `json:"file,omitempty" bencode:"file,omitempty"`
	Code                   string                 `json:"code,omitempty" bencode:"code,omitempty"`
	DTLSFingerprint        DTLSFingerprint        `json:"DTLS-fingerprint,omitempty" bencode:"DTLS-fingerprint,omitempty"`
	ICELite                string                 `json:"ICE-lite,omitempty" bencode:"ICE-lite,omitempty"`
	MediaEcho              string                 `json:"media-echo,omitempty" bencode:"media-echo,omitempty"`
	Label                  string                 `json:"label,omitempty" bencode:"label,omitempty"`
	SetLabel               string                 `json:"set-label,omitempty" bencode:"set-label,omitempty"`
	FromLabel              string                 `json:"from-label,omitempty" bencode:"from-label,omitempty"`
	ToLabel                string                 `json:"to-label,omitempty" bencode:"to-label,omitempty"`
	DTMFSecurity           string                 `json:"DTMF-security,omitempty" bencode:"DTMF-security,omitempty"`
	Digit                  string                 `json:"digit,omitempty" bencode:"digit,omitempty"`
	DTMFSecurityTrigger    string                 `json:"DTMF-security-trigger,omitempty" bencode:"DTMF-security-trigger,omitempty"`
	DTMFSecurityTriggerEnd string                 `json:"DTMF-security-trigger-end,omitempty" bencode:"DTMF-security-trigger-end,omitempty"`
	Trigger                string                 `json:"trigger,omitempty" bencode:"trigger,omitempty"`
	TriggerEnd             string                 `json:"trigger-end,omitempty" bencode:"trigger-end,omitempty"`
	All                    string                 `json:"all,omitempty" bencode:"all,omitempty"`
	Frequency              string                 `json:"frequency,omitempty" bencode:"frequency,omitempty"`
	Blob                   string                 `json:"blob,omitempty" bencode:"blob,omitempty"`
	Sdp                    string                 `json:"sdp,omitempty" bencode:"sdp,omitempty"`
	AudioPlayer            string                 `json:"audio-player,omitempty" bencode:"audio-player,omitempty"`
	DTMFLogDest            string                 `json:"dtmf-log-dest,omitempty" bencode:"dtmf-log-dest,omitempty"`
	OutputDestination      string                 `json:"output-destination,omitempty" bencode:"output-destination,omitempty"`
	VscStartRec            string                 `json:"vsc-start-rec,omitempty" bencode:"vsc-start-rec,omitempty"`
	VscStopRec             string                 `json:"vsc-stop-rec,omitempty" bencode:"vsc-stop-rec,omitempty"`
	VscPauseRec            string                 `json:"vsc-pause-rec,omitempty" bencode:"vsc-pause-rec,omitempty"`
	VscStartStopRec        string                 `json:"vsc-start-stop-rec,omitempty" bencode:"vsc-start-stop-rec,omitempty"`
	VscPauseResumeRec      string                 `json:"vsc-pause-resume-rec,omitempty" bencode:"vsc-pause-resume-rec,omitempty"`
	VscStartPauseResumeRec string                 `json:"vsc-start-pause-resume-rec,omitempty" bencode:"vsc-start-pause-resume-rec,omitempty"`
	RtppFlags              string                 `json:"rtpp-flags,omitempty" bencode:"rtpp-flags,omitempty"`
	SdpAttr                *ParamsSdpAttrSections `json:"sdp-attr,omitempty" bencode:"sdp-attr,omitempty"`
}

// Parametros de comportamento tipo inteiro
type ParamsOptInt struct {
	TOS              int `json:"TOS,omitempty" bencode:"TOS,omitempty"`
	DeleteDelay      int `json:"delete-delay,omitempty" bencode:"delete-delay,omitempty"`
	DelayBuffer      int `json:"delay-buffer,omitempty" bencode:"delay-buffer,omitempty"`
	Volume           int `json:"volume,omitempty" bencode:"volume,omitempty"`
	TriggerEndTime   int `json:"trigger-end-time,omitempty" bencode:"trigger-end-time,omitempty"`
	TriggerEndDigits int `json:"trigger-end-digits,omitempty" bencode:"trigger-end-digits,omitempty"`
	DTMFDelay        int `json:"DTMF-delay,omitempty" bencode:"DTMF-delay,omitempty"`
	Ptime            int `json:"ptime,omitempty" bencode:"ptime,omitempty"`
	PtimeReverse     int `json:"ptime-reverse,omitempty" bencode:"ptime-reverse,omitempty"`
	DbId             int `json:"db-id,omitempty" bencode:"db-id,omitempty"`
	Duration         int `json:"duration,omitempty" bencode:"duration,omitempty"`
}

// Parametros de comportamento tipo array separado por ','
type ParamsOptStringArray struct {
	Flags        []ParamFlags   `json:"flags,omitempty" bencode:"flags,omitempty"`
	RtcpMux      []ParamRTCPMux `json:"rtcp-mux,omitempty" bencode:"rtcp-mux,omitempty"`
	SDES         []SDES         `json:"SDES,omitempty" bencode:"SDES,omitempty"`
	Supports     []string       `json:"supports,omitempty" bencode:"supports,omitempty"`
	T38          []string       `json:"T38,omitempty" bencode:"T38,omitempty"`
	OSRTP        []OSRTP        `json:"OSRTP,omitempty" bencode:"OSRTP,omitempty"`
	ReceivedFrom []string       `json:"received-from,omitempty" bencode:"received-from,omitempty"`
	FromTags     []string       `json:"from-tags,omitempty" bencode:"from-tags,omitempty"`
	Frequencies  []string       `json:"frequencies,omitempty" bencode:"frequencies,omitempty"`
	Replace      []ParamReplace `json:"replace,omitempty" bencode:"replace,omitempty"`
}

// Parametros de manipulação de sessão
type ParamsSdpAttrSections struct {
	Global *ParamsSdpAttrCommands `json:"global,omitempty" bencode:"global,omitempty"`
	Audio  *ParamsSdpAttrCommands `json:"audio,omitempty" bencode:"audio,omitempty"`
	Video  *ParamsSdpAttrCommands `json:"video,omitempty" bencode:"video,omitempty"`
	None   *ParamsSdpAttrCommands `json:"none,omitempty" bencode:"none,omitempty"`
}

// Parametros de atributos de comandos
type ParamsSdpAttrCommands struct {
	Add        []string   `json:"add,omitempty" bencode:"add,omitempty"`
	Remove     []string   `json:"remove,omitempty" bencode:"remove,omitempty"`
	Substitute [][]string `json:"substitute,omitempty" bencode:"substitute,omitempty"`
}

// Gera o cookie do comando
func (r *Engine) GetCookie() string {
	return uuid.NewString()
}

// Atribuir o ip padrão para conexão
func (r *Engine) GetIP() net.IP {
	return r.ip
}

// Atribuir a porta padrão para conexão
func (r *Engine) GetPort() int {
	return r.port
}

// Atribuir o protocolo padrão para conexão
func (r *Engine) GetProto() string {
	return r.proto
}

// Atribuir a porta padrão NG porta de controler
func (r *Engine) GetNG() int {
	return r.ng
}

// Abrir conexão com o proxy rtpengine
func (r *Engine) Conn() (net.Conn, error) {
	engine := r.ip.String() + ":" + fmt.Sprint(r.port)
	conn, err := net.Dial(r.proto, engine)

	if err != nil {
		fmt.Println(err.Error(), r.proto, engine)
		return nil, err
	}

	defer net.Dial(r.proto, engine)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	r.con = conn
	return r.con, nil

}

// Abrir conexão com o proxy rtpengine
func (r *Engine) ConnUDP() (*net.UDPConn, error) {
	engine := r.ip.String() + ":" + fmt.Sprint(r.port)
	addr := &net.UDPAddr{
		IP:   r.ip,
		Port: r.port,
	}
	conn, err := net.DialUDP(r.proto, nil, addr)

	if err != nil {
		fmt.Println(err.Error(), r.proto, engine)
		return nil, err
	}

	defer net.DialUDP(r.proto, nil, addr)
	conn.SetReadDeadline(time.Now().Add(time.Minute))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	r.conUDP = conn
	return r.conUDP, nil

}

// Trasformar o comando em bencode
func EncodeComando(cookie string, command *RequestRtp) ([]byte, error) {
	data, err := bencode.Marshal(command)
	if err != nil {
		return nil, err
	}

	bind := []byte(cookie + " ")
	return append(bind, data...), nil
}

func DecodeResposta(cookie string, resposta []byte) *ResponseRtp {
	resp := &ResponseRtp{}
	cookieIndex := bytes.IndexAny(resposta, " ")
	if cookieIndex != len(cookie) {
		resp.Result = "error"
		resp.ErrorReason = "Erro ao analisar a mensagem"
		return resp
	}

	cookieResponse := string(resposta[:cookieIndex])
	if cookieResponse != cookie {
		resp.Result = "error"
		resp.ErrorReason = "O cookie não corresponde"
		return resp
	}

	encodedData := string(resposta[cookieIndex+1:])
	decodedDataRaw, err := ben.NewParserFromString(encodedData).AsDict()
	if err != nil {
		return resp
	}

	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   resp,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(decodedDataRaw)

	return resp
}
