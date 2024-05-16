package rtpengine

// Definição do Protocolo de Transporte do SDP
type TransportProtocol string

const (
	RTP_AVP           TransportProtocol = "RTP/AVP"
	RTP_SAVP          TransportProtocol = "RTP/SAVP"
	RTP_AVPF          TransportProtocol = "RTP/AVPF"
	RTP_SAVPF         TransportProtocol = "RTP/SAVPF"
	UDP_TLS_RTP_SAVP  TransportProtocol = "UDP/TLS/RTP/SAVP"
	UDP_TLS_RTP_SAVPF TransportProtocol = "UDP/TLS/RTP/SAVPF"
)

// Definição dos comandos aceitos

type TipoComandos string

const (
	Ping             TipoComandos = "ping"
	Offer            TipoComandos = "offer"
	Answer           TipoComandos = "answer"
	Delete           TipoComandos = "delete"
	Query            TipoComandos = "query"
	List             TipoComandos = "list"
	StartRecording   TipoComandos = "start recording"
	StopRecording    TipoComandos = "stop recording"
	PauseRecording   TipoComandos = "pause recording"
	BlockDTMF        TipoComandos = "block DTMF"
	UnblockDTMF      TipoComandos = "unblock DTMF"
	BlockMedia       TipoComandos = "block media"
	UnblockMedia     TipoComandos = "unblock media"
	SilenceMedia     TipoComandos = "silence media"
	UnsilenceMedia   TipoComandos = "unsilence media"
	StartForwarding  TipoComandos = "start forwarding"
	StopForwarding   TipoComandos = "stop forwarding"
	PlayMedia        TipoComandos = "play media"
	StopMedia        TipoComandos = "stop media"
	PlayDTMF         TipoComandos = "play DTMF"
	Statistics       TipoComandos = "statistics"
	Publish          TipoComandos = "publish"
	SubscribeRequest TipoComandos = "subscribe request"
	SubscribeAnswer  TipoComandos = "subscribe answer"
	Unsubscribe      TipoComandos = "unsubscribe"
)

// Definição dos tipo dtls
type DtlsHash string

const (
	Sha256 DtlsHash = "sha-256"
	Sha1   DtlsHash = "sha-1"
	Sha224 DtlsHash = "sha-224"
	Sha384 DtlsHash = "sha-384"
	Sha512 DtlsHash = "sha-512"
)

// Tipo de parametros para o CryptoSuite
type CryptoSuite string

const (
	SRTP_AEAD_AES_256_GCM         CryptoSuite = "AEAD_AES_256_GCM"
	SRTP_AEAD_AES_128_GCM         CryptoSuite = "AEAD_AES_128_GCM"
	SRTP_AES_256_CM_HMAC_SHA1_80  CryptoSuite = "AES_256_CM_HMAC_SHA1_80"
	SRTP_AES_256_CM_HMAC_SHA1_32  CryptoSuite = "AES_256_CM_HMAC_SHA1_32"
	SRTP_AES_192_CM_HMAC_SHA1_80  CryptoSuite = "AES_192_CM_HMAC_SHA1_80"
	SRTP_AES_192_CM_HMAC_SHA1_32  CryptoSuite = "AES_192_CM_HMAC_SHA1_32"
	SRTP_AES_CM_128_HMAC_SHA1_80  CryptoSuite = "AES_CM_128_HMAC_SHA1_80"
	SRTP_AAES_CM_128_HMAC_SHA1_32 CryptoSuite = "AES_CM_128_HMAC_SHA1_32"
	SRTP_F8_128_HMAC_SHA1_80      CryptoSuite = "F8_128_HMAC_SHA1_80"
	SRTP_F8_128_HMAC_SHA1_32      CryptoSuite = "F8_128_HMAC_SHA1_32"
	SRTP_NULL_HMAC_SHA1_80        CryptoSuite = "NULL_HMAC_SHA1_80"
	SRTP_NULL_HMAC_SHA1_32        CryptoSuite = "NULL_HMAC_SHA1_32"
)

// Tipo de parametros para o replace
type ParamReplace string

const (
	Origin                   ParamReplace = "origin"
	SessionConnection        ParamReplace = "session-connection"
	SdpVersion               ParamReplace = "sdp-version"
	Username                 ParamReplace = "username"
	SessionName              ParamReplace = "session-name"
	ZeroAddress              ParamReplace = "zero-address"
	ForceIncrementSdpVersion ParamReplace = "force-increment-sdp-version"
)

// Tipo de parametros usado como flags
type ParamFlags string

const (
	TrustAddress         ParamFlags = "trust-address"
	Symmetric            ParamFlags = "symmetric"
	Asymmetric           ParamFlags = "asymmetric"
	Unidirectional       ParamFlags = "unidirectional"
	Force                ParamFlags = "force"
	StrictSource         ParamFlags = "strict-source"
	MediaHandover        ParamFlags = "media-handover"
	SipSourceAddress     ParamFlags = "sip-source-address"
	Reset                ParamFlags = "reset"
	PortLatching         ParamFlags = "port-latching"
	NoRtcpAttribute      ParamFlags = "no-rtcp-attribute"
	FullRtcpAttribute    ParamFlags = "full-rtcp-attribute"
	LoopProtect          ParamFlags = "loop-protect"
	RecordCall           ParamFlags = "record-call"
	AlwaysTranscode      ParamFlags = "always-transcode"
	SIPREC               ParamFlags = "SIPREC"
	PadCrypto            ParamFlags = "pad-crypto"
	GenerateMid          ParamFlags = "generate-mid"
	Fragment             ParamFlags = "fragment"
	OriginalSendrecv     ParamFlags = "original-sendrecv"
	SymmetricCodecs      ParamFlags = "symmetric-codecs"
	AsymmetricCodecs     ParamFlags = "asymmetric-codecs"
	InjectDTMF           ParamFlags = "inject-DTMF"
	DetectDTMF           ParamFlags = "detect-DTMF"
	GenerateRTCP         ParamFlags = "generate-RTCP"
	SingleCodec          ParamFlags = "single-codec"
	NoCodecRenegotiation ParamFlags = "no-codec-renegotiation"
	PierceNAT            ParamFlags = "pierce-NAT"
	SIPSourceAddress     ParamFlags = "SIP-source-address"
	AllowTranscoding     ParamFlags = "allow-transcoding"
	TrickleICE           ParamFlags = "trickle-ICE"
	RejectICE            ParamFlags = "reject-ICE"
	Egress               ParamFlags = "egress"
	NoJitterBuffer       ParamFlags = "no-jitter-buffer"
	Passthrough          ParamFlags = "passthrough"
	NoPassthrough        ParamFlags = "no-passthrough"
	Pause                ParamFlags = "pause"
	EarlyMedia           ParamFlags = "early-media"
	BlockShort           ParamFlags = "block-short"
	RecordingVsc         ParamFlags = "recording-vsc"
	BlockEgress          ParamFlags = "block-egress"
	StripExtmap          ParamFlags = "strip-extmap"
)

// Tipo de parametros usado no rtcp-mux
type ParamRTCPMux string

const (
	RTCP_Offer   ParamRTCPMux = "offer"
	RTCP_Require ParamRTCPMux = "require"
	RTCP_Demux   ParamRTCPMux = "demux"
	RTCP_Accept  ParamRTCPMux = "accept"
	RTCP_Reject  ParamRTCPMux = "reject"
)

// Tipo de string codecs
type Codecs string

const (
	CODEC_PCMU  Codecs = "PCMU"
	CODEC_PCMA  Codecs = "PCMA"
	CODEC_G729  Codecs = "G729"
	CODEC_G729a Codecs = "G729a"
	CODEC_OPUS  Codecs = "opus"
	CODEC_G722  Codecs = "G722"
	CODEC_G723  Codecs = "G723"
	CODEC_ILBC  Codecs = "iLBC"
	CODEC_SPEEX Codecs = "speex"
)

// Tipo de string ICE
type ICE string

const (
	ICERemove     ICE = "remove"
	ICEForce      ICE = "force"
	ICEDefault    ICE = "default"
	ICEForceRelay ICE = "force-relay"
	ICEOptional   ICE = "optional"
)

// Tipo de string DTLS
type DTLS string

const (
	DTLSOff     DTLS = "off"
	DTLSNo      DTLS = "no"
	DTLSDisable DTLS = "disable"
	DTLSPassive DTLS = "passive"
	DTLSActive  DTLS = "active"
)

// Tipo DTLS reverso string
type DTLSReverse string

const (
	DTLSReversePassive DTLSReverse = "passive"
	DTLSReverseActive  DTLSReverse = "active"
)

// Tipo DTLS-fingerprint string
type DTLSFingerprint string

const (
	DTLSFingerprintSha256 DTLSFingerprint = "sha-256"
	DTLSFingerprintSha1   DTLSFingerprint = "sha-1"
	DTLSFingerprintSha224 DTLSFingerprint = "sha-224"
	DTLSFingerprintSha384 DTLSFingerprint = "sha-384"
	DTLSFingerprintSha512 DTLSFingerprint = "sha-512"
)
