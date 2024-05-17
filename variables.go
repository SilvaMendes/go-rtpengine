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
	SdpVersion               ParamReplace = "SDP-version"
	Username                 ParamReplace = "username"
	SessionName              ParamReplace = "session-name"
	ZeroAddress              ParamReplace = "zero-address"
	ForceIncrementSdpVersion ParamReplace = "force-increment-sdp-version"
	ForceIncrementSdpVer     ParamReplace = "force-increment-sdp-ver"
)

// Tipo de parametros usado como flags
type ParamFlags string

const (
	TrustAddress          ParamFlags = "trust-address"
	Symmetric             ParamFlags = "symmetric"
	Asymmetric            ParamFlags = "asymmetric"
	Unidirectional        ParamFlags = "unidirectional"
	Force                 ParamFlags = "force"
	StrictSource          ParamFlags = "strict-source"
	MediaHandover         ParamFlags = "media-handover"
	Reset                 ParamFlags = "reset"
	PortLatching          ParamFlags = "port-latching"
	NoRtcpAttribute       ParamFlags = "no-rtcp-attribute"
	FullRtcpAttribute     ParamFlags = "full-rtcp-attribute"
	LoopProtect           ParamFlags = "loop-protect"
	RecordCall            ParamFlags = "record-call"
	AlwaysTranscode       ParamFlags = "always-transcode"
	SIPREC                ParamFlags = "SIPREC"
	PadCrypto             ParamFlags = "pad-crypto"
	GenerateMid           ParamFlags = "generate-mid"
	Fragment              ParamFlags = "fragment"
	OriginalSendrecv      ParamFlags = "original-sendrecv"
	SymmetricCodecs       ParamFlags = "symmetric-codecs"
	AsymmetricCodecs      ParamFlags = "asymmetric-codecs"
	InjectDTMF            ParamFlags = "inject-DTMF"
	DetectDTMF            ParamFlags = "detect-DTMF"
	GenerateRTCP          ParamFlags = "generate-RTCP"
	SingleCodec           ParamFlags = "single-codec"
	NoCodecRenegotiation  ParamFlags = "no-codec-renegotiation"
	PierceNAT             ParamFlags = "pierce-NAT"
	SIPSourceAddress      ParamFlags = "SIP-source-address"
	AllowTranscoding      ParamFlags = "allow-transcoding"
	TrickleICE            ParamFlags = "trickle-ICE"
	RejectICE             ParamFlags = "reject-ICE"
	Egress                ParamFlags = "egress"
	NoJitterBuffer        ParamFlags = "no-jitter-buffer"
	Passthrough           ParamFlags = "passthrough"
	NoPassthrough         ParamFlags = "no-passthrough"
	Pause                 ParamFlags = "pause"
	EarlyMedia            ParamFlags = "early-media"
	BlockShort            ParamFlags = "block-short"
	RecordingVsc          ParamFlags = "recording-vsc"
	BlockEgress           ParamFlags = "block-egress"
	StripExtmap           ParamFlags = "strip-extmap"
	NATWait               ParamFlags = "NAT-wait"
	NoPortLatching        ParamFlags = "no-port-latching"
	RecordingAnnouncement ParamFlags = "recording-announcement"
	ReuseCodecs           ParamFlags = "reuse-codecs"
	RTCPMirror            ParamFlags = "RTCP-mirror"
	StaticCodecs          ParamFlags = "static-codecs"
)

// Tipo de parametros usado no rtcp-mux
type ParamRTCPMux string

const (
	RTCPOffer   ParamRTCPMux = "offer"
	RTCPRequire ParamRTCPMux = "require"
	RTCPDemux   ParamRTCPMux = "demux"
	RTCPAccept  ParamRTCPMux = "accept"
	RTCPReject  ParamRTCPMux = "reject"
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

// Tipo SDES string
type SDES string

const (
	SDESOff                         SDES = "off"
	SDESNo                          SDES = "no"
	SDESDisable                     SDES = "disable"
	SDESNonew                       SDES = "nonew"
	SDESPad                         SDES = "pad"
	SDESStatic                      SDES = "static"
	SDESPrefer                      SDES = "prefer"
	SDESUnencrypted_srtp            SDES = "unencrypted_srtp"
	SDESUnencrypted_srtcp           SDES = "unencrypted_srtcp"
	SDESUnauthenticated_srtp        SDES = "unauthenticated_srtp"
	SDESEncrypted_srtp              SDES = "encrypted_srtp"
	SDESEncrypted_srtcp             SDES = "encrypted_srtcp"
	SDESAuthenticated_srtp          SDES = "authenticated_srtp"
	SDESNoAEAD_AES_256_GCM          SDES = "no-AEAD_AES_256_GCM"
	SDESNoAEAD_AES_128_GCM          SDES = "no-AEAD_AES_128_GCM"
	SDESNoAES_256_CM_HMAC_SHA1_80   SDES = "no-AES_256_CM_HMAC_SHA1_80"
	SDESNoAES_256_CM_HMAC_SHA1_32   SDES = "no-AES_256_CM_HMAC_SHA1_32"
	SDESNoAES_192_CM_HMAC_SHA1_80   SDES = "no-AES_192_CM_HMAC_SHA1_80"
	SDESNoAES_192_CM_HMAC_SHA1_32   SDES = "no-AES_192_CM_HMAC_SHA1_32"
	SDESNoAES_CM_128_HMAC_SHA1_80   SDES = "no-AES_CM_128_HMAC_SHA1_80"
	SDESNoAES_CM_128_HMAC_SHA1_32   SDES = "no-AES_CM_128_HMAC_SHA1_32"
	SDESNoF8_128_HMAC_SHA1_80       SDES = "no-F8_128_HMAC_SHA1_80"
	SDESNoF8_128_HMAC_SHA1_32       SDES = "no-F8_128_HMAC_SHA1_32"
	SDESNoNULL_HMAC_SHA1_80         SDES = "no-NULL_HMAC_SHA1_80"
	SDESNoNULL_HMAC_SHA1_32         SDES = "no-NULL_HMAC_SHA1_32"
	SDESOnlyAEAD_AES_256_GCM        SDES = "only-AEAD_AES_256_GCM"
	SDESOnlyAEAD_AES_128_GCM        SDES = "only-AEAD_AES_128_GCM"
	SDESOnlyAES_256_CM_HMAC_SHA1_80 SDES = "only-AES_256_CM_HMAC_SHA1_80"
	SDESOnlyAES_256_CM_HMAC_SHA1_32 SDES = "only-AES_256_CM_HMAC_SHA1_32"
	SDESOnlyAES_192_CM_HMAC_SHA1_80 SDES = "only-AES_192_CM_HMAC_SHA1_80"
	SDESOnlyAES_192_CM_HMAC_SHA1_32 SDES = "only-AES_192_CM_HMAC_SHA1_32"
	SDESOnlyAES_CM_128_HMAC_SHA1_80 SDES = "only-AES_CM_128_HMAC_SHA1_80"
	SDESOnlyAES_CM_128_HMAC_SHA1_32 SDES = "only-AES_CM_128_HMAC_SHA1_32"
	SDESOnlyF8_128_HMAC_SHA1_80     SDES = "only-F8_128_HMAC_SHA1_80"
	SDESOnlyF8_128_HMAC_SHA1_32     SDES = "only-F8_128_HMAC_SHA1_32"
	SDESOnlyNULL_HMAC_SHA1_80       SDES = "only-NULL_HMAC_SHA1_80"
	SDESOnlyNULL_HMAC_SHA1_32       SDES = "only-NULL_HMAC_SHA1_32"
)

// Tipo OSRTP string
type OSRTP string

const (
	OSRTPOffer        OSRTP = "offer"
	OSRTPOfferRFC     OSRTP = "offer-RFC"
	OSRTPOfferLegacy  OSRTP = "offer-legacy"
	OSRTPAcceptRFC    OSRTP = "accept-RFC"
	OSRTPAcceptLegacy OSRTP = "accept-legacy"
	OSRTPAccept       OSRTP = "accept"
)
