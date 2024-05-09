package rtpengine

// Definição do Protocolo de Transporte do SDP
type TransportProtocol string

// Definição dos comandos aceitos

type TipoComandos string

// Definição dos tipo dtls
type DtlsHash string

type CryptoSuite string

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
)

// Estrutura da requisicão do comando
type RequestRtp struct {
	Command struct{} `json:"command" bencode:"command"`
}

// Estrutura da resposta do comando
type ResponseRtp struct {
	Result struct{} `json:"result" bencode:"result"`
}

// Parametros usado como passagem de flags
type ParamsFlags struct {
	TrustAddress         string `json:"trust-address,omitempty" bencode:"trust-address,omitempty"`
	Symmetric            string `json:"symmetric,omitempty" bencode:"symmetric,omitempty"`
	Asymmetric           string `json:"asymmetric,omitempty" bencode:"asymmetric,omitempty"`
	Unidirectional       string `json:"unidirectional,omitempty" bencode:"asymmetric,omitempty"`
	Force                string `json:"force,omitempty" bencode:"force,omitempty"`
	StrictSource         string `json:"strict-source,omitempty" bencode:"strict-source,omitempty"`
	MediaHandover        string `json:"media-handover,omitempty" bencode:"media-handover,omitempty"`
	SipSourceAddress     string `json:"sip-source-address,omitempty" bencode:"sip-source-address,omitempty"`
	Reset                string `json:"reset,omitempty" bencode:"reset,omitempty"`
	PortLatching         string `json:"port-latching,omitempty" bencode:"port-latching,omitempty"`
	NoRtcpAttribute      string `json:"no-rtcp-attribute,omitempty" bencode:"no-rtcp-attribute,omitempty" `
	FullRtcpAttribute    string `json:"full-rtcp-attribute,omitempty" bencode:"full-rtcp-attribute,omitempty"`
	LoopProtect          string `json:"loop-protect,omitempty" bencode:"loop-protect,omitempty"`
	RecordCall           string `json:"record-call,omitempty" bencode:"record-call,omitempty"`
	AlwaysTranscode      string `json:"always-transcode,omitempty" bencode:"always-transcode,omitempty"`
	SIPREC               string `json:"SIPREC,omitempty" bencode:"SIPREC,omitempty"`
	PadCrypto            string `json:"pad-crypto,omitempty" bencode:"pad-crypto,omitempty"`
	GenerateMid          string `json:"generate-mid,omitempty" bencode:"generate-mid,omitempty"`
	Fragment             string `json:"fragment,omitempty" bencode:"fragment,omitempty"`
	OriginalSendrecv     string `json:"original-sendrecv,omitempty" bencode:"original-sendrecv,omitempty"`
	SymmetricCodecs      string `json:"symmetric-codecs,omitempty" bencode:"symmetric-codecs,omitempty"`
	AsymmetricCodecs     string `json:"asymmetric-codecs,omitempty" bencode:"asymmetric-codecs,omitempty"`
	InjectDTMF           string `json:"inject-DTMF,omitempty" bencode:"asymmetric-codecs,omitempty"`
	DetectDTMF           string `json:"detect-DTMF,omitempty" bencode:"detect-DTMF,omitempty"`
	GenerateRTCP         string `json:"generate-RTCP,omitempty" bencode:"generate-RTCP,omitempty"`
	SingleCodec          string `json:"single-codec,omitempty" bencode:"single-codec,omitempty"`
	NoCodecRenegotiation string `json:"no-codec-renegotiation,omitempty" bencode:"no-codec-renegotiation,omitempty"`
	PierceNAT            string `json:"pierce-NAT,omitempty" bencode:"pierce-NAT,omitempty"`
	SIPSourceAddress     string `json:"SIP-source-address,omitempty" bencode:"SIP-source-address,omitempty"`
	AllowTranscoding     string `json:"allow-transcoding,omitempty" bencode:"allow-transcoding,omitempty"`
	TrickleICE           string `json:"trickle-ICE,omitempty" bencode:"trickle-ICE,omitempty"`
	RejectICE            string `json:"reject-ICE,omitempty" bencode:"reject-ICE,omitempty"`
	Egress               string `json:"egress,omitempty" bencode:"egress,omitempty"`
	NoJitterBuffer       string `json:"no-jitter-buffer,omitempty" bencode:"no-jitter-buffer,omitempty"`
	Passthrough          string `json:"passthrough,omitempty" bencode:"passthrough,omitempty"`
	NoPassthrough        string `json:"no-passthrough,omitempty" bencode:"no-passthrough,omitempty"`
	Pause                string `json:"pause,omitempty" bencode:"pause,omitempty" `
	EarlyMedia           string `json:"early-media,omitempty" bencode:"early-media,omitempty"`
	BlockShort           string `json:"block-short,omitempty" bencode:"block-short,omitempty"`
	RecordingVsc         string `json:"recording-vsc,omitempty" bencode:"recording-vsc,omitempty"`
	BlockEgress          string `json:"block-egress,omitempty" bencode:"block-egress,omitempty"`
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
	Flags        string `json:"flags,omitempty" bencode:"flags,omitempty"`
	RtcpMux      string `json:"rtcp-mux,omitempty" bencode:"rtcp-mux,omitempty"`
	SDES         string `json:"SDES,omitempty" bencode:"SDES,omitempty"`
	Supports     string `json:"supports,omitempty" bencode:"supports,omitempty"`
	T38          string `json:"T38,omitempty" bencode:"T38,omitempty"`
	OSRTP        string `json:"OSRTP,omitempty" bencode:"OSRTP,omitempty"`
	ReceivedFrom string `json:"received-from,omitempty" bencode:"received-from,omitempty"`
	FromTags     string `json:"from-tags,omitempty" bencode:"from-tags,omitempty"`
	Frequencies  string `json:"frequencies,omitempty" bencode:"frequencies,omitempty"`
}

// Parametros de manipulação dos codecs na oferta
type ParamsOptCodec struct {
	Strip     string `json:"strip,omitempty" bencode:"strip,omitempty"`
	Offer     string `json:"offer,omitempty" bencode:"offer,omitempty"`
	Transcode string `json:"transcode,omitempty" bencode:"transcode,omitempty"`
	Mask      string `json:"mask,omitempty" bencode:"mask,omitempty"`
	Set       string `json:"set,omitempty" bencode:"set,omitempty"`
	Consume   string `json:"consume,omitempty" bencode:"consume,omitempty"`
	Accept    string `json:"accept,omitempty" bencode:"accept,omitempty"`
	Except    string `json:"except,omitempty" bencode:"except,omitempty"`
}

// Parametros de substituição de valores no corpo do SDP
type ParamsOptReplace struct {
	Origin                   string `json:"origin,omitempty" bencode:"origin,omitempty"`
	SessionConnection        string `json:"session-connection,omitempty" bencode:"session-connection,omitempty"`
	SdpVersion               string `json:"sdp-version,omitempty" bencode:"sdp-version,omitempty"`
	Username                 string `json:"username,omitempty" bencode:"username,omitempty"`
	SessionName              string `json:"session-name,omitempty" bencode:"session-name,omitempty"`
	ZeroAddress              string `json:"zero-address,omitempty" bencode:"zero-address,omitempty"`
	ForceIncrementSdpVersion string `json:"force-increment-sdp-version,omitempty" bencode:"force-increment-sdp-version,omitempty"`
}

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
