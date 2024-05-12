package rtpengine

import (
	"bytes"
	"fmt"
	"net"

	bencode "github.com/anacrolix/torrent/bencode"
	"github.com/google/uuid"
)

type Engine struct {
	con   net.Conn
	ip    net.IP
	port  int
	dns   *net.Resolver
	proto string
	ng    int
}

// Definição do Protocolo de Transporte do SDP
type TransportProtocol string

// Definição dos comandos aceitos

type TipoComandos string

// Definição dos tipo dtls
type DtlsHash string

type CryptoSuite string

// Tipo de parametros para o replace
type ParamReplace string

// Tipo de parametros usado como flags
type ParamFlags string

// Tipo de parametros usado no rtcp-mux
type ParamRTCPMux string

//// Tipo de codec para transcoder
//type TranscoderCodec string

//// Tipo de Mask para remover codec
//type MaskCodec string

// Tipo de codecs
type Codecs string

const (
	Ping                          TipoComandos      = "ping"
	Offer                         TipoComandos      = "offer"
	Answer                        TipoComandos      = "answer"
	Delete                        TipoComandos      = "delete"
	Query                         TipoComandos      = "query"
	List                          TipoComandos      = "list"
	StartRecording                TipoComandos      = "start recording"
	StopRecording                 TipoComandos      = "stop recording"
	PauseRecording                TipoComandos      = "pause recording"
	BlockDTMF                     TipoComandos      = "block DTMF"
	UnblockDTMF                   TipoComandos      = "unblock DTMF"
	BlockMedia                    TipoComandos      = "block media"
	UnblockMedia                  TipoComandos      = "unblock media"
	SilenceMedia                  TipoComandos      = "silence media"
	UnsilenceMedia                TipoComandos      = "unsilence media"
	StartForwarding               TipoComandos      = "start forwarding"
	StopForwarding                TipoComandos      = "stop forwarding"
	PlayMedia                     TipoComandos      = "play media"
	StopMedia                     TipoComandos      = "stop media"
	PlayDTMF                      TipoComandos      = "play DTMF"
	Statistics                    TipoComandos      = "statistics"
	Publish                       TipoComandos      = "publish"
	SubscribeRequest              TipoComandos      = "subscribe request"
	SubscribeAnswer               TipoComandos      = "subscribe answer"
	Unsubscribe                   TipoComandos      = "unsubscribe"
	RTP_AVP                       TransportProtocol = "RTP/AVP"
	RTP_SAVP                      TransportProtocol = "RTP/SAVP"
	RTP_AVPF                      TransportProtocol = "RTP/AVPF"
	RTP_SAVPF                     TransportProtocol = "RTP/SAVPF"
	UDP_TLS_RTP_SAVP              TransportProtocol = "UDP/TLS/RTP/SAVP"
	UDP_TLS_RTP_SAVPF             TransportProtocol = "UDP/TLS/RTP/SAVPF"
	Sha256                        DtlsHash          = "sha-256"
	Sha1                          DtlsHash          = "sha-1"
	Sha224                        DtlsHash          = "sha-224"
	Sha384                        DtlsHash          = "sha-384"
	Sha512                        DtlsHash          = "sha-512"
	SRTP_AEAD_AES_256_GCM         CryptoSuite       = "AEAD_AES_256_GCM"
	SRTP_AEAD_AES_128_GCM         CryptoSuite       = "AEAD_AES_128_GCM"
	SRTP_AES_256_CM_HMAC_SHA1_80  CryptoSuite       = "AES_256_CM_HMAC_SHA1_80"
	SRTP_AES_256_CM_HMAC_SHA1_32  CryptoSuite       = "AES_256_CM_HMAC_SHA1_32"
	SRTP_AES_192_CM_HMAC_SHA1_80  CryptoSuite       = "AES_192_CM_HMAC_SHA1_80"
	SRTP_AES_192_CM_HMAC_SHA1_32  CryptoSuite       = "AES_192_CM_HMAC_SHA1_32"
	SRTP_AES_CM_128_HMAC_SHA1_80  CryptoSuite       = "AES_CM_128_HMAC_SHA1_80"
	SRTP_AAES_CM_128_HMAC_SHA1_32 CryptoSuite       = "AES_CM_128_HMAC_SHA1_32"
	SRTP_F8_128_HMAC_SHA1_80      CryptoSuite       = "F8_128_HMAC_SHA1_80"
	SRTP_F8_128_HMAC_SHA1_32      CryptoSuite       = "F8_128_HMAC_SHA1_32"
	SRTP_NULL_HMAC_SHA1_80        CryptoSuite       = "NULL_HMAC_SHA1_80"
	SRTP_NULL_HMAC_SHA1_32        CryptoSuite       = "NULL_HMAC_SHA1_32"
	Origin                        ParamReplace      = "origin"
	SessionConnection             ParamReplace      = "session-connection"
	SdpVersion                    ParamReplace      = "sdp-version"
	Username                      ParamReplace      = "username"
	SessionName                   ParamReplace      = "session-name"
	ZeroAddress                   ParamReplace      = "zero-address"
	ForceIncrementSdpVersion      ParamReplace      = "force-increment-sdp-version"
	TrustAddress                  ParamFlags        = "trust-address"
	Symmetric                     ParamFlags        = "symmetric"
	Asymmetric                    ParamFlags        = "asymmetric"
	Unidirectional                ParamFlags        = "unidirectional"
	Force                         ParamFlags        = "force"
	StrictSource                  ParamFlags        = "strict-source"
	MediaHandover                 ParamFlags        = "media-handover"
	SipSourceAddress              ParamFlags        = "sip-source-address"
	Reset                         ParamFlags        = "reset"
	PortLatching                  ParamFlags        = "port-latching"
	NoRtcpAttribute               ParamFlags        = "no-rtcp-attribute"
	FullRtcpAttribute             ParamFlags        = "full-rtcp-attribute"
	LoopProtect                   ParamFlags        = "loop-protect"
	RecordCall                    ParamFlags        = "record-call"
	AlwaysTranscode               ParamFlags        = "always-transcode"
	SIPREC                        ParamFlags        = "SIPREC"
	PadCrypto                     ParamFlags        = "pad-crypto"
	GenerateMid                   ParamFlags        = "generate-mid"
	Fragment                      ParamFlags        = "fragment"
	OriginalSendrecv              ParamFlags        = "original-sendrecv"
	SymmetricCodecs               ParamFlags        = "symmetric-codecs"
	AsymmetricCodecs              ParamFlags        = "asymmetric-codecs"
	InjectDTMF                    ParamFlags        = "inject-DTMF"
	DetectDTMF                    ParamFlags        = "detect-DTMF"
	GenerateRTCP                  ParamFlags        = "generate-RTCP"
	SingleCodec                   ParamFlags        = "single-codec"
	NoCodecRenegotiation          ParamFlags        = "no-codec-renegotiation"
	PierceNAT                     ParamFlags        = "pierce-NAT"
	SIPSourceAddress              ParamFlags        = "SIP-source-address"
	AllowTranscoding              ParamFlags        = "allow-transcoding"
	TrickleICE                    ParamFlags        = "trickle-ICE"
	RejectICE                     ParamFlags        = "reject-ICE"
	Egress                        ParamFlags        = "egress"
	NoJitterBuffer                ParamFlags        = "no-jitter-buffer"
	Passthrough                   ParamFlags        = "passthrough"
	NoPassthrough                 ParamFlags        = "no-passthrough"
	Pause                         ParamFlags        = "pause"
	EarlyMedia                    ParamFlags        = "early-media"
	BlockShort                    ParamFlags        = "block-short"
	RecordingVsc                  ParamFlags        = "recording-vsc"
	BlockEgress                   ParamFlags        = "block-egress"
	RTCP_Offer                    ParamRTCPMux      = "offer"
	RTCP_Require                  ParamRTCPMux      = "require"
	RTCP_Demux                    ParamRTCPMux      = "demux"
	RTCP_Accept                   ParamRTCPMux      = "accept"
	RTCP_Reject                   ParamRTCPMux      = "reject"
	CODEC_PCMU                    Codecs            = "PCMU"
	CODEC_PCMA                    Codecs            = "PCMA"
	CODEC_G729                    Codecs            = "G729"
	CODEC_G729a                   Codecs            = "G729a"
	CODEC_OPUS                    Codecs            = "opus"
	CODEC_G722                    Codecs            = "G722"
	CODEC_G723                    Codecs            = "G723"
	CODEC_ILBC                    Codecs            = "iLBC"
	CODEC_SPEEX                   Codecs            = "speex"
)

// Estrutura da requisicão do comando
type RequestRtp struct {
	Command string `json:"command" bencode:"command"`
	*ParamsOptString
	*ParamsOptInt
	*ParamsOptStringArray
	//*ParamsOptCodec
	*ParamsSdpAttrSections
	*ParamsSdpAttrCommands
}

// Estrutura da resposta do comando
type ResponseRtp struct {
	Result      string      `json:"result" bencode:"result"`
	Sdp         string      `json:"sdp,omitempty" bencode:"sdp,omitempty"`
	ErrorReason string      `json:"error-reason,omitempty" bencode:"error-reason,omitempty"`
	Warning     string      `json:"warning,omitempty" bencode:"warning,omitempty"`
	Created     int         `json:"created,omitempty" bencode:"created,omitempty"`
	CreatedUs   int         `json:"created_us,omitempty" bencode:"created_us,omitempty"`
	LastSignal  int         `json:"last signal,omitempty" bencode:"last signal,omitempty"`
	SSRC        interface{} `json:"SSRC,omitempty" bencode:"SSRC,omitempty"`
	Tags        interface{} `json:"tags,omitempty" bencode:"tags,omitempty"`
	Totals      TotalRTP    `json:"totals,omitempty" bencode:"totals,omitempty"`
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
	FromTag                string `json:"from-tag" bencode:"from-tag"`
	ToTag                  string `json:"to-tag" bencode:"to-tag"`
	CallId                 string `json:"call-id" bencode:"call-id"`
	TransportProtocol      string `json:"transport-protocol" bencode:"transport-protocol"`
	MediaAddress           string `json:"media-address,omitempty" bencode:"media-address,omitempty"`
	ICE                    string `json:"ICE,omitempty" bencode:"ICE,omitempty"`
	AddressFamily          string `json:"address-family,omitempty" bencode:"address-family,omitempty"`
	DTLS                   string `json:"DTLS,omitempty" bencode:"DTLS,omitempty"`
	ViaBranch              string `json:"via-branch,omitempty" bencode:"via-branch,omitempty"`
	XmlrpcCallback         string `json:"xmlrpc-callback,omitempty" bencode:"xmlrpc-callback,omitempty"`
	Metadata               string `json:"metadata,omitempty" bencode:"metadata,omitempty"`
	Address                string `json:"address,omitempty" bencode:"address,omitempty"`
	File                   string `json:"file,omitempty" bencode:"file,omitempty"`
	Code                   string `json:"code,omitempty" bencode:"code,omitempty"`
	DTLSFingerprint        string `json:"DTLS-fingerprint,omitempty" bencode:"DTLS-fingerprint,omitempty"`
	ICELite                string `json:"ICE-lite,omitempty" bencode:"ICE-lite,omitempty"`
	MediaEcho              string `json:"media-echo,omitempty" bencode:"media-echo,omitempty"`
	Label                  string `json:"label,omitempty" bencode:"label,omitempty"`
	SetLabel               string `json:"set-label,omitempty" bencode:"set-label,omitempty"`
	FromLabel              string `json:"from-label,omitempty" bencode:"from-label,omitempty"`
	ToLabel                string `json:"to-label,omitempty" bencode:"to-label,omitempty"`
	DTMFSecurity           string `json:"DTMF-security,omitempty" bencode:"DTMF-security,omitempty"`
	Digit                  string `json:"digit,omitempty" bencode:"digit,omitempty"`
	DTMFSecurityTrigger    string `json:"DTMF-security-trigger,omitempty" bencode:"DTMF-security-trigger,omitempty"`
	DTMFSecurityTriggerEnd string `json:"DTMF-security-trigger-end,omitempty" bencode:"DTMF-security-trigger-end,omitempty"`
	Trigger                string `json:"trigger,omitempty" bencode:"trigger,omitempty"`
	TriggerEnd             string `json:"trigger-end,omitempty" bencode:"trigger-end,omitempty"`
	All                    string `json:"all,omitempty" bencode:"all,omitempty"`
	Frequency              string `json:"frequency,omitempty" bencode:"frequency,omitempty"`
	Blob                   string `json:"blob,omitempty" bencode:"blob,omitempty"`
	Sdp                    string `json:"sdp" bencode:"sdp"`
	AudioPlayer            string `json:"audio-player,omitempty" bencode:"audio-player,omitempty"`
	DTMFLogDest            string `json:"dtmf-log-dest,omitempty" bencode:"dtmf-log-dest,omitempty"`
	OutputDestination      string `json:"output-destination,omitempty" bencode:"output-destination,omitempty"`
	VscStartRec            string `json:"vsc-start-rec,omitempty" bencode:"vsc-start-rec,omitempty"`
	VscStopRec             string `json:"vsc-stop-rec,omitempty" bencode:"vsc-stop-rec,omitempty"`
	VscPauseRec            string `json:"vsc-pause-rec,omitempty" bencode:"vsc-pause-rec,omitempty"`
	VscStartStopRec        string `json:"vsc-start-stop-rec,omitempty" bencode:"vsc-start-stop-rec,omitempty"`
	VscPauseResumeRec      string `json:"vsc-pause-resume-rec,omitempty" bencode:"vsc-pause-resume-rec,omitempty"`
	VscStartPauseResumeRec string `json:"vsc-start-pause-resume-rec,omitempty" bencode:"vsc-start-pause-resume-rec,omitempty"`
	RtppFlags              string `json:"rtpp-flags,omitempty" bencode:"rtpp-flags,omitempty"`
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
	DbId             int `json:"db-id,omitempty" bencode:"db-id,omitempty"`
	Duration         int `json:"duration,omitempty" bencode:"duration,omitempty"`
}

// Parametros de comportamento tipo array separado por ','
type ParamsOptStringArray struct {
	Flags        []string       `json:"flags,omitempty" bencode:"flags,omitempty"`
	RtcpMux      []ParamRTCPMux `json:"rtcp-mux,omitempty" bencode:"rtcp-mux,omitempty"`
	SDES         []string       `json:"SDES,omitempty" bencode:"SDES,omitempty"`
	Supports     []string       `json:"supports,omitempty" bencode:"supports,omitempty"`
	T38          []string       `json:"T38,omitempty" bencode:"T38,omitempty"`
	OSRTP        []string       `json:"OSRTP,omitempty" bencode:"OSRTP,omitempty"`
	ReceivedFrom []string       `json:"received-from,omitempty" bencode:"received-from,omitempty"`
	FromTags     []string       `json:"from-tags,omitempty" bencode:"from-tags,omitempty"`
	Frequencies  []string       `json:"frequencies,omitempty" bencode:"frequencies,omitempty"`
	Replace      []ParamReplace `json:"replace,omitempty" bencode:"replace,omitempty"`
}

//// Parametros de manipulação dos codecs na oferta
//type ParamsOptCodec struct {
//	Strip     string   `json:"strip,omitempty" bencode:"strip,omitempty"`
//	Offer     string   `json:"offer,omitempty" bencode:"offer,omitempty"`
//	Transcode []string `json:"transcode,omitempty" bencode:"transcode,omitempty"`
//	Mask      string   `json:"mask,omitempty" bencode:"mask,omitempty"`
//	Set       string   `json:"set,omitempty" bencode:"set,omitempty"`
//	Consume   string   `json:"consume,omitempty" bencode:"consume,omitempty"`
//	Accept    string   `json:"accept,omitempty" bencode:"accept,omitempty"`
//	Except    string   `json:"except,omitempty" bencode:"except,omitempty"`
//}

// Parametros de manipulação de sessão
type ParamsSdpAttrSections struct {
	Global string `json:"global,omitempty" bencode:"global,omitempty"`
	Audio  string `json:"audio,omitempty" bencode:"audio,omitempty"`
	Video  string `json:"video,omitempty" bencode:"video,omitempty"`
}

// Parametros de atributos de comandos
type ParamsSdpAttrCommands struct {
	Add    string `json:"add,omitempty" bencode:"add,omitempty"`
	Remove string `json:"remove,omitempty" bencode:"remove,omitempty"`
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

// Abrir conexão com p proxy rtpengine
func (r *Engine) Conn() (net.Conn, error) {
	engine := r.ip.String() + ":" + fmt.Sprint(r.port)
	conn, err := net.Dial(r.proto, engine)
	if err != nil {
		fmt.Println(err.Error(), r.proto, engine)
		return nil, err
	}
	r.con = conn
	return r.con, nil

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
	err := bencode.Unmarshal([]byte(encodedData), resp)

	if err != nil {
		return resp
	}

	return resp
}
