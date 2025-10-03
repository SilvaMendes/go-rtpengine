/*
Package rtpengine

This file defines several types and constants used for RTP (Real-time Transport Protocol) session management and configuration.

Types and Constants:

- TransportProtocol: Supported transport protocols (e.g., RTP/AVP, RTP/SAVP).
- TypeCommands: Possible commands for RTP session control (e.g., offer, answer, delete).
- DtlsHash: Supported DTLS hash algorithms.
- CryptoSuite: Supported SRTP crypto suites.
- ParamReplace: SDP session parameters that can be replaced.
- ParamFlags: Flags that modify RTP session behavior.
- ParamRTCPMux: RTCP multiplexing options.
- Codecs: Supported audio codecs.
- ICE: ICE negotiation options.
- DTLS: DTLS operation modes.
- DTLSReverse: Reverse DTLS modes.
- DTLSFingerprint: DTLS fingerprint algorithms.
- SDES: SDES negotiation options.
- OSRTP: OSRTP operation modes.
- AddressFamily: Supported address families (IPv4/IPv6).
- Connection: Connection types (e.g., zero for music on hold).
- Record: Call recording options.

Each type is defined as a string and has constants representing valid values for use in RTP engine operations.
*/

package rtpengine

type TransportProtocol string

const (
	RTP_AVP           TransportProtocol = "RTP/AVP"
	RTP_SAVP          TransportProtocol = "RTP/SAVP"
	RTP_AVPF          TransportProtocol = "RTP/AVPF"
	RTP_SAVPF         TransportProtocol = "RTP/SAVPF"
	UDP_TLS_RTP_SAVP  TransportProtocol = "UDP/TLS/RTP/SAVP"
	UDP_TLS_RTP_SAVPF TransportProtocol = "UDP/TLS/RTP/SAVPF"
)

type TypeCommands string

const (
	Ping             TypeCommands = "ping"
	Offer            TypeCommands = "offer"
	Answer           TypeCommands = "answer"
	Delete           TypeCommands = "delete"
	Query            TypeCommands = "query"
	List             TypeCommands = "list"
	StartRecording   TypeCommands = "start recording"
	StopRecording    TypeCommands = "stop recording"
	PauseRecording   TypeCommands = "pause recording"
	BlockDTMF        TypeCommands = "block DTMF"
	UnblockDTMF      TypeCommands = "unblock DTMF"
	BlockMedia       TypeCommands = "block media"
	UnblockMedia     TypeCommands = "unblock media"
	SilenceMedia     TypeCommands = "silence media"
	UnsilenceMedia   TypeCommands = "unsilence media"
	StartForwarding  TypeCommands = "start forwarding"
	StopForwarding   TypeCommands = "stop forwarding"
	PlayMedia        TypeCommands = "play media"
	StopMedia        TypeCommands = "stop media"
	PlayDTMF         TypeCommands = "play DTMF"
	Statistics       TypeCommands = "statistics"
	Publish          TypeCommands = "publish"
	SubscribeRequest TypeCommands = "subscribe request"
	SubscribeAnswer  TypeCommands = "subscribe answer"
	Unsubscribe      TypeCommands = "unsubscribe"
	Connect          TypeCommands = "connect"
)

type DtlsHash string

const (
	Sha256 DtlsHash = "sha-256"
	Sha1   DtlsHash = "sha-1"
	Sha224 DtlsHash = "sha-224"
	Sha384 DtlsHash = "sha-384"
	Sha512 DtlsHash = "sha-512"
)

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

type ParamReplace string

const (
	// DEPRECADO replace-session-connection flag encountered, but not supported anymore
	Origin                   ParamReplace = "origin"
	OriginFull               ParamReplace = "origin-full"
	SessionConnection        ParamReplace = "session-connection"
	SdpVersion               ParamReplace = "SDP-version"
	Username                 ParamReplace = "username"
	SessionName              ParamReplace = "session-name"
	ZeroAddress              ParamReplace = "zero-address"
	ForceIncrementSdpVersion ParamReplace = "force-increment-sdp-version"
	ForceIncrementSdpVer     ParamReplace = "force-increment-sdp-ver"
)

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
	CodecExceptPCMU       ParamFlags = "codec-except-PCMU"
	CodecExceptPCMA       ParamFlags = "codec-except-PCMA"
	CodecExceptG729       ParamFlags = "codec-except-G729"
	CodecExceptG729a      ParamFlags = "codec-except-G729a"
	CodecExceptOpus       ParamFlags = "codec-except-opus"
	CodecExceptG722       ParamFlags = "codec-except-G722"
	CodecExceptG723       ParamFlags = "codec-except-G723"
	CodecExceptILBC       ParamFlags = "codec-except-iLBC"
	CodecExceptSpeex      ParamFlags = "codec-except-speex"
	CodecStripPCMU        ParamFlags = "codec-strip-PCMU"
	CodecStripPCMA        ParamFlags = "codec-strip-PCMA"
	CodecStripG729        ParamFlags = "codec-strip-G729"
	CodecStripG729a       ParamFlags = "codec-strip-G729a"
	CodecStripOpus        ParamFlags = "codec-strip-opus"
	CodecStripG722        ParamFlags = "codec-strip-G722"
	CodecStripG723        ParamFlags = "codec-strip-G723"
	CodecStripILBC        ParamFlags = "codec-strip-iLBC"
	CodecStripSpeex       ParamFlags = "codec-strip-speex"
	CodecMaskPCMA         ParamFlags = "codec-mask-PCMA"
	CodecMaskG729         ParamFlags = "codec-mask-G729"
	CodecMaskG729a        ParamFlags = "codec-mask-G729a"
	CodecMaskOpus         ParamFlags = "codec-mask-opus"
	CodecMaskG722         ParamFlags = "codec-mask-G722"
	CodecMaskG723         ParamFlags = "codec-mask-G723"
	CodecMaskILBC         ParamFlags = "codec-mask-iLBC"
	CodecMaskSpeex        ParamFlags = "codec-mask-speex"
	CodecTranscodePCMA    ParamFlags = "codec-transcode-PCMA"
	CodecTranscodeG729    ParamFlags = "codec-transcode-G729"
	CodecTranscodeG729a   ParamFlags = "codec-transcode-G729a"
	CodecTranscodeOpus    ParamFlags = "codec-transcode-opus"
	CodecTranscodeG722    ParamFlags = "codec-transcode-G722"
	CodecTranscodeG723    ParamFlags = "codec-transcode-G723"
	CodecTranscodeILBC    ParamFlags = "codec-transcode-iLBC"
	CodecTranscodeSpeex   ParamFlags = "codec-transcode-speex"
)

type ParamRTCPMux string

const (
	RTCPOffer   ParamRTCPMux = "offer"
	RTCPRequire ParamRTCPMux = "require"
	RTCPDemux   ParamRTCPMux = "demux"
	RTCPAccept  ParamRTCPMux = "accept"
	RTCPReject  ParamRTCPMux = "reject"
)

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

type ICE string

const (
	ICERemove     ICE = "remove"
	ICEForce      ICE = "force"
	ICEDefault    ICE = "default"
	ICEForceRelay ICE = "force-relay"
	ICEOptional   ICE = "optional"
)

type DTLS string

const (
	DTLSOff     DTLS = "off"
	DTLSNo      DTLS = "no"
	DTLSDisable DTLS = "disable"
	DTLSPassive DTLS = "passive"
	DTLSActive  DTLS = "active"
)

type DTLSReverse string

const (
	DTLSReversePassive DTLSReverse = "passive"
	DTLSReverseActive  DTLSReverse = "active"
)

type DTLSFingerprint string

const (
	DTLSFingerprintSha256 DTLSFingerprint = "sha-256"
	DTLSFingerprintSha1   DTLSFingerprint = "sha-1"
	DTLSFingerprintSha224 DTLSFingerprint = "sha-224"
	DTLSFingerprintSha384 DTLSFingerprint = "sha-384"
	DTLSFingerprintSha512 DTLSFingerprint = "sha-512"
)

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

type OSRTP string

const (
	OSRTPOffer        OSRTP = "offer"
	OSRTPOfferRFC     OSRTP = "offer-RFC"
	OSRTPOfferLegacy  OSRTP = "offer-legacy"
	OSRTPAcceptRFC    OSRTP = "accept-RFC"
	OSRTPAcceptLegacy OSRTP = "accept-legacy"
	OSRTPAccept       OSRTP = "accept"
)

type AddressFamily string

const (
	AddressFamilyIP4 AddressFamily = "IP4"
	AddressFamilyIP6 AddressFamily = "IP6"
)

type Connection string

const (
	MohConnection Connection = "zero"
)

type Record string

const (
	RecordYes Record = "yes"
	RecordNo  Record = "no"
	RecordOn  Record = "on"
	RecordOff Record = "off"
)
