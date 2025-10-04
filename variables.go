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

// TransportProtocol defines the supported transport protocols for RTP communication.
type TransportProtocol string

const (
	// RTP_AVP represents the standard RTP Audio/Video Profile.
	RTP_AVP TransportProtocol = "RTP/AVP"

	// RTP_SAVP represents the Secure RTP Audio/Video Profile using SRTP.
	RTP_SAVP TransportProtocol = "RTP/SAVP"

	// RTP_AVPF represents the RTP Audio/Video Profile with Feedback.
	RTP_AVPF TransportProtocol = "RTP/AVPF"

	// RTP_SAVPF represents the Secure RTP Audio/Video Profile with Feedback using SRTP.
	RTP_SAVPF TransportProtocol = "RTP/SAVPF"

	// UDP_TLS_RTP_SAVP represents RTP over UDP with TLS encryption using the Secure Audio/Video Profile.
	UDP_TLS_RTP_SAVP TransportProtocol = "UDP/TLS/RTP/SAVP"

	// UDP_TLS_RTP_SAVPF represents RTP over UDP with TLS encryption using the Secure Audio/Video Profile with Feedback.
	UDP_TLS_RTP_SAVPF TransportProtocol = "UDP/TLS/RTP/SAVPF"
)

// TypeCommands defines the set of supported command types for RTP engine operations.
type TypeCommands string

const (
	// Ping is used to check the availability or responsiveness of the RTP engine.
	Ping TypeCommands = "ping"

	// Offer initiates a media session by sending an SDP offer to the RTP engine.
	Offer TypeCommands = "offer"

	// Answer responds to an SDP offer to establish a media session.
	Answer TypeCommands = "answer"

	// Delete terminates an existing media session.
	Delete TypeCommands = "delete"

	// Query retrieves information about an existing media session.
	Query TypeCommands = "query"

	// List returns a list of all active media sessions.
	List TypeCommands = "list"

	// StartRecording begins recording the media stream of a session.
	StartRecording TypeCommands = "start recording"

	// StopRecording stops the ongoing media recording of a session.
	StopRecording TypeCommands = "stop recording"

	// PauseRecording temporarily pauses the media recording of a session.
	PauseRecording TypeCommands = "pause recording"

	// BlockDTMF blocks DTMF (Dual-tone multi-frequency) signals in a media session.
	BlockDTMF TypeCommands = "block DTMF"

	// UnblockDTMF allows DTMF signals in a media session.
	UnblockDTMF TypeCommands = "unblock DTMF"

	// BlockMedia blocks media transmission in a session.
	BlockMedia TypeCommands = "block media"

	// UnblockMedia resumes media transmission in a session.
	UnblockMedia TypeCommands = "unblock media"

	// SilenceMedia mutes the media stream in a session.
	SilenceMedia TypeCommands = "silence media"

	// UnsilenceMedia unmutes the media stream in a session.
	UnsilenceMedia TypeCommands = "unsilence media"

	// StartForwarding begins forwarding media to a specified destination.
	StartForwarding TypeCommands = "start forwarding"

	// StopForwarding stops forwarding media to a specified destination.
	StopForwarding TypeCommands = "stop forwarding"

	// PlayMedia plays a media file into the session.
	PlayMedia TypeCommands = "play media"

	// StopMedia stops the playback of a media file in the session.
	StopMedia TypeCommands = "stop media"

	// PlayDTMF sends a DTMF tone into the media session.
	PlayDTMF TypeCommands = "play DTMF"

	// Statistics retrieves performance and usage statistics from the RTP engine.
	Statistics TypeCommands = "statistics"

	// Publish sends media stream information to the RTP engine for distribution.
	Publish TypeCommands = "publish"

	// SubscribeRequest initiates a subscription request for media stream updates.
	SubscribeRequest TypeCommands = "subscribe request"

	// SubscribeAnswer responds to a subscription request for media stream updates.
	SubscribeAnswer TypeCommands = "subscribe answer"

	// Unsubscribe cancels an existing media stream subscription.
	Unsubscribe TypeCommands = "unsubscribe"

	// Connect establishes a connection between media endpoints.
	Connect TypeCommands = "connect"
)

// DtlsHash defines the supported hash algorithms used in DTLS (Datagram Transport Layer Security).
type DtlsHash string

const (
	// Sha256 represents the SHA-256 hash algorithm used in DTLS.
	Sha256 DtlsHash = "sha-256"

	// Sha1 represents the SHA-1 hash algorithm used in DTLS.
	Sha1 DtlsHash = "sha-1"

	// Sha224 represents the SHA-224 hash algorithm used in DTLS.
	Sha224 DtlsHash = "sha-224"

	// Sha384 represents the SHA-384 hash algorithm used in DTLS.
	Sha384 DtlsHash = "sha-384"

	// Sha512 represents the SHA-512 hash algorithm used in DTLS.
	Sha512 DtlsHash = "sha-512"
)

// CryptoSuite defines the supported SRTP (Secure Real-time Transport Protocol) cryptographic suites.
// These suites specify the encryption and authentication algorithms used to secure RTP media streams.
type CryptoSuite string

const (
	// SRTP_AEAD_AES_256_GCM uses AES-256 in Galois/Counter Mode (GCM) for encryption and authentication.
	SRTP_AEAD_AES_256_GCM CryptoSuite = "AEAD_AES_256_GCM"

	// SRTP_AEAD_AES_128_GCM uses AES-128 in Galois/Counter Mode (GCM) for encryption and authentication.
	SRTP_AEAD_AES_128_GCM CryptoSuite = "AEAD_AES_128_GCM"

	// SRTP_AES_256_CM_HMAC_SHA1_80 uses AES-256 in Counter Mode (CM) with HMAC-SHA1 authentication (80-bit).
	SRTP_AES_256_CM_HMAC_SHA1_80 CryptoSuite = "AES_256_CM_HMAC_SHA1_80"

	// SRTP_AES_256_CM_HMAC_SHA1_32 uses AES-256 in Counter Mode (CM) with HMAC-SHA1 authentication (32-bit).
	SRTP_AES_256_CM_HMAC_SHA1_32 CryptoSuite = "AES_256_CM_HMAC_SHA1_32"

	// SRTP_AES_192_CM_HMAC_SHA1_80 uses AES-192 in Counter Mode (CM) with HMAC-SHA1 authentication (80-bit).
	SRTP_AES_192_CM_HMAC_SHA1_80 CryptoSuite = "AES_192_CM_HMAC_SHA1_80"

	// SRTP_AES_192_CM_HMAC_SHA1_32 uses AES-192 in Counter Mode (CM) with HMAC-SHA1 authentication (32-bit).
	SRTP_AES_192_CM_HMAC_SHA1_32 CryptoSuite = "AES_192_CM_HMAC_SHA1_32"

	// SRTP_AES_CM_128_HMAC_SHA1_80 uses AES-128 in Counter Mode (CM) with HMAC-SHA1 authentication (80-bit).
	SRTP_AES_CM_128_HMAC_SHA1_80 CryptoSuite = "AES_CM_128_HMAC_SHA1_80"

	// SRTP_AAES_CM_128_HMAC_SHA1_32 uses AES-128 in Counter Mode (CM) with HMAC-SHA1 authentication (32-bit).
	SRTP_AAES_CM_128_HMAC_SHA1_32 CryptoSuite = "AES_CM_128_HMAC_SHA1_32"

	// SRTP_F8_128_HMAC_SHA1_80 uses AES-128 in F8 Mode with HMAC-SHA1 authentication (80-bit).
	SRTP_F8_128_HMAC_SHA1_80 CryptoSuite = "F8_128_HMAC_SHA1_80"

	// SRTP_F8_128_HMAC_SHA1_32 uses AES-128 in F8 Mode with HMAC-SHA1 authentication (32-bit).
	SRTP_F8_128_HMAC_SHA1_32 CryptoSuite = "F8_128_HMAC_SHA1_32"

	// SRTP_NULL_HMAC_SHA1_80 uses no encryption and HMAC-SHA1 authentication (80-bit).
	SRTP_NULL_HMAC_SHA1_80 CryptoSuite = "NULL_HMAC_SHA1_80"

	// SRTP_NULL_HMAC_SHA1_32 uses no encryption and HMAC-SHA1 authentication (32-bit).
	SRTP_NULL_HMAC_SHA1_32 CryptoSuite = "NULL_HMAC_SHA1_32"
)

// ParamReplace defines the types of SDP (Session Description Protocol) parameters
// that can be modified or replaced during RTP engine operations.
type ParamReplace string

const (
	// Origin replaces the SDP origin field.
	Origin ParamReplace = "origin"

	// OriginFull replaces the full SDP origin line, including username, session ID, version, network type, address type, and address.
	OriginFull ParamReplace = "origin-full"

	// SessionConnection replaces the SDP session-level connection information.
	// Deprecated: The 'replace-session-connection' flag is no longer supported.
	SessionConnection ParamReplace = "session-connection"

	// SdpVersion replaces the SDP version field.
	SdpVersion ParamReplace = "SDP-version"

	// Username replaces the username field in the SDP origin line.
	Username ParamReplace = "username"

	// SessionName replaces the session name field in the SDP.
	SessionName ParamReplace = "session-name"

	// ZeroAddress replaces the connection address with a zero IP address (e.g., 0.0.0.0).
	ZeroAddress ParamReplace = "zero-address"

	// ForceIncrementSdpVersion forces an increment of the SDP version number.
	ForceIncrementSdpVersion ParamReplace = "force-increment-sdp-version"

	// ForceIncrementSdpVer is an alias for force-increment-sdp-version.
	ForceIncrementSdpVer ParamReplace = "force-increment-sdp-ver"
)

// ParamFlags defines a set of optional flags that modify the behavior of RTP engine operations.
// These flags control aspects of media handling, signaling, codec negotiation, NAT traversal, and recording.
type ParamFlags string

const (
	// TrustAddress trusts the source IP address from the signaling layer.
	TrustAddress ParamFlags = "trust-address"

	// Symmetric enables symmetric RTP behavior (same IP/port for sending and receiving).
	Symmetric ParamFlags = "symmetric"

	// Asymmetric enables asymmetric RTP behavior (different IP/port for sending and receiving).
	Asymmetric ParamFlags = "asymmetric"

	// Unidirectional forces one-way media flow.
	Unidirectional ParamFlags = "unidirectional"

	// Force applies the operation even if conditions are not ideal.
	Force ParamFlags = "force"

	// StrictSource enforces strict source IP verification for incoming media.
	StrictSource ParamFlags = "strict-source"

	// MediaHandover enables media handover between endpoints.
	MediaHandover ParamFlags = "media-handover"

	// Reset resets the media session state.
	Reset ParamFlags = "reset"

	// PortLatching enables port latching for NAT traversal.
	PortLatching ParamFlags = "port-latching"

	// NoRtcpAttribute disables RTCP attribute in SDP.
	NoRtcpAttribute ParamFlags = "no-rtcp-attribute"

	// FullRtcpAttribute includes full RTCP attribute in SDP.
	FullRtcpAttribute ParamFlags = "full-rtcp-attribute"

	// LoopProtect prevents media loopback scenarios.
	LoopProtect ParamFlags = "loop-protect"

	// RecordCall enables call recording.
	RecordCall ParamFlags = "record-call"

	// AlwaysTranscode forces media transcoding regardless of codec compatibility.
	AlwaysTranscode ParamFlags = "always-transcode"

	// SIPREC enables SIP Recording (SIPREC) support.
	SIPREC ParamFlags = "SIPREC"

	// PadCrypto pads crypto attributes in SDP.
	PadCrypto ParamFlags = "pad-crypto"

	// GenerateMid generates media stream identifiers (MID) in SDP.
	GenerateMid ParamFlags = "generate-mid"

	// Fragment enables SDP fragmentation.
	Fragment ParamFlags = "fragment"

	// OriginalSendrecv preserves original sendrecv attributes.
	OriginalSendrecv ParamFlags = "original-sendrecv"

	// SymmetricCodecs enforces symmetric codec negotiation.
	SymmetricCodecs ParamFlags = "symmetric-codecs"

	// AsymmetricCodecs allows asymmetric codec negotiation.
	AsymmetricCodecs ParamFlags = "asymmetric-codecs"

	// InjectDTMF injects DTMF tones into the media stream.
	InjectDTMF ParamFlags = "inject-DTMF"

	// DetectDTMF enables detection of DTMF tones.
	DetectDTMF ParamFlags = "detect-DTMF"

	// GenerateRTCP enables RTCP packet generation.
	GenerateRTCP ParamFlags = "generate-RTCP"

	// SingleCodec restricts media to a single codec.
	SingleCodec ParamFlags = "single-codec"

	// NoCodecRenegotiation disables codec renegotiation during session updates.
	NoCodecRenegotiation ParamFlags = "no-codec-renegotiation"

	// PierceNAT attempts to pierce NAT for media flow.
	PierceNAT ParamFlags = "pierce-NAT"

	// SIPSourceAddress uses the SIP source address for media routing.
	SIPSourceAddress ParamFlags = "SIP-source-address"

	// AllowTranscoding permits media transcoding.
	AllowTranscoding ParamFlags = "allow-transcoding"

	// TrickleICE enables Trickle ICE support.
	TrickleICE ParamFlags = "trickle-ICE"

	// RejectICE disables ICE candidate negotiation.
	RejectICE ParamFlags = "reject-ICE"

	// Egress marks the media flow as outbound.
	Egress ParamFlags = "egress"

	// NoJitterBuffer disables jitter buffering.
	NoJitterBuffer ParamFlags = "no-jitter-buffer"

	// Passthrough allows media to pass through without modification.
	Passthrough ParamFlags = "passthrough"

	// NoPassthrough disables media passthrough.
	NoPassthrough ParamFlags = "no-passthrough"

	// Pause temporarily halts media transmission.
	Pause ParamFlags = "pause"

	// EarlyMedia enables early media handling before call is answered.
	EarlyMedia ParamFlags = "early-media"

	// BlockShort blocks short-duration media packets.
	BlockShort ParamFlags = "block-short"

	// RecordingVsc enables VSC (Video Surveillance Control) recording.
	RecordingVsc ParamFlags = "recording-vsc"

	// BlockEgress blocks outbound media flow.
	BlockEgress ParamFlags = "block-egress"

	// StripExtmap removes extmap attributes from SDP.
	StripExtmap ParamFlags = "strip-extmap"

	// NATWait waits for NAT traversal before media transmission.
	NATWait ParamFlags = "NAT-wait"

	// NoPortLatching disables port latching.
	NoPortLatching ParamFlags = "no-port-latching"

	// RecordingAnnouncement enables recording announcements.
	RecordingAnnouncement ParamFlags = "recording-announcement"

	// ReuseCodecs reuses previously negotiated codecs.
	ReuseCodecs ParamFlags = "reuse-codecs"

	// RTCPMirror mirrors RTCP packets.
	RTCPMirror ParamFlags = "RTCP-mirror"

	// StaticCodecs enforces use of static codec payload types.
	StaticCodecs ParamFlags = "static-codecs"

	// CodecExceptPCMU excludes PCMU codec from negotiation.
	CodecExceptPCMU ParamFlags = "codec-except-PCMU"

	// CodecExceptPCMA excludes PCMA codec from negotiation.
	CodecExceptPCMA ParamFlags = "codec-except-PCMA"

	// CodecExceptG729 excludes G729 codec from negotiation.
	CodecExceptG729 ParamFlags = "codec-except-G729"

	// CodecExceptG729a excludes G729a codec from negotiation.
	CodecExceptG729a ParamFlags = "codec-except-G729a"

	// CodecExceptOpus excludes Opus codec from negotiation.
	CodecExceptOpus ParamFlags = "codec-except-opus"

	// CodecExceptG722 excludes G722 codec from negotiation.
	CodecExceptG722 ParamFlags = "codec-except-G722"

	// CodecExceptG723 excludes G723 codec from negotiation.
	CodecExceptG723 ParamFlags = "codec-except-G723"

	// CodecExceptILBC excludes iLBC codec from negotiation.
	CodecExceptILBC ParamFlags = "codec-except-iLBC"

	// CodecExceptSpeex excludes Speex codec from negotiation.
	CodecExceptSpeex ParamFlags = "codec-except-speex"

	// CodecStripPCMU removes PCMU codec from SDP.
	CodecStripPCMU ParamFlags = "codec-strip-PCMU"

	// CodecStripPCMA removes PCMA codec from SDP.
	CodecStripPCMA ParamFlags = "codec-strip-PCMA"

	// CodecStripG729 removes G729 codec from SDP.
	CodecStripG729 ParamFlags = "codec-strip-G729"

	// CodecStripG729a removes G729a codec from SDP.
	CodecStripG729a ParamFlags = "codec-strip-G729a"

	// CodecStripOpus removes Opus codec from SDP.
	CodecStripOpus ParamFlags = "codec-strip-opus"

	// CodecStripG722 removes G722 codec from SDP.
	CodecStripG722 ParamFlags = "codec-strip-G722"

	// CodecStripG723 removes G723 codec from SDP.
	CodecStripG723 ParamFlags = "codec-strip-G723"

	// CodecStripILBC removes iLBC codec from SDP.
	CodecStripILBC ParamFlags = "codec-strip-iLBC"

	// CodecStripSpeex removes Speex codec from SDP.
	CodecStripSpeex ParamFlags = "codec-strip-speex"

	// CodecMaskPCMA masks PCMA codec in SDP.
	CodecMaskPCMA ParamFlags = "codec-mask-PCMA"

	// CodecMaskG729 masks G729 codec in SDP.
	CodecMaskG729 ParamFlags = "codec-mask-G729"

	// CodecMaskG729a masks G729a codec in SDP.
	CodecMaskG729a ParamFlags = "codec-mask-G729a"

	// CodecMaskOpus masks Opus codec in SDP.
	CodecMaskOpus ParamFlags = "codec-mask-opus"

	// CodecMaskG722 masks G722 codec in SDP.
	CodecMaskG722 ParamFlags = "codec-mask-G722"

	// CodecMaskG723 masks G723 codec in SDP.
	CodecMaskG723 ParamFlags = "codec-mask-G723"

	// CodecMaskILBC masks iLBC codec in SDP.
	CodecMaskILBC ParamFlags = "codec-mask-iLBC"

	// CodecMaskSpeex masks Speex codec in SDP.
	CodecMaskSpeex ParamFlags = "codec-mask-speex"

	// CodecTranscodePCMA enables transcoding to PCMA codec.
	CodecTranscodePCMA ParamFlags = "codec-transcode-PCMA"

	// CodecTranscodeG729 enables transcoding to G729 codec.
	CodecTranscodeG729 ParamFlags = "codec-transcode-G729"

	// CodecTranscodeG729a enables transcoding to G729a codec.
	CodecTranscodeG729a ParamFlags = "codec-transcode-G729a"

	// CodecTranscodeOpus enables transcoding to Opus codec.
	CodecTranscodeOpus ParamFlags = "codec-transcode-opus"

	// CodecTranscodeG722 enables transcoding to G722 codec.
	CodecTranscodeG722 ParamFlags = "codec-transcode-G722"

	// CodecTranscodeG723 enables transcoding to G723 codec.
	CodecTranscodeG723 ParamFlags = "codec-transcode-G723"

	// CodecTranscodeILBC enables transcoding to iLBC codec.
	CodecTranscodeILBC ParamFlags = "codec-transcode-iLBC"

	// CodecTranscodeSpeex enables transcoding to Speex codec.
	CodecTranscodeSpeex ParamFlags = "codec-transcode-speex"
)

// ParamRTCPMux defines the RTCP multiplexing modes used in RTP engine operations.
// RTCP multiplexing allows RTP and RTCP packets to be sent over a single port, improving efficiency and NAT traversal.
type ParamRTCPMux string

const (
	// RTCPOffer indicates that RTCP multiplexing is being offered.
	RTCPOffer ParamRTCPMux = "offer"

	// RTCPRequire indicates that RTCP multiplexing is required and must be supported.
	RTCPRequire ParamRTCPMux = "require"

	// RTCPDemux enables demultiplexing of RTP and RTCP packets received on the same port.
	RTCPDemux ParamRTCPMux = "demux"

	// RTCPAccept accepts RTCP multiplexing if offered by the remote party.
	RTCPAccept ParamRTCPMux = "accept"

	// RTCPReject rejects RTCP multiplexing, forcing separate RTP and RTCP streams.
	RTCPReject ParamRTCPMux = "reject"
)

// Codecs defines the supported audio codecs for RTP media sessions.
// These codecs determine the format and compression used for audio transmission.
type Codecs string

const (
	// CODEC_PCMU represents the G.711 µ-law codec, commonly used in North America and Japan.
	CODEC_PCMU Codecs = "PCMU"

	// CODEC_PCMA represents the G.711 A-law codec, commonly used in Europe and international systems.
	CODEC_PCMA Codecs = "PCMA"

	// CODEC_G729 represents the G.729 codec, a low-bitrate codec suitable for VoIP applications.
	CODEC_G729 Codecs = "G729"

	// CODEC_G729a represents the G.729a codec, a variant of G.729 with reduced complexity.
	CODEC_G729a Codecs = "G729a"

	// CODEC_OPUS represents the Opus codec, a versatile and high-quality codec for both voice and audio.
	CODEC_OPUS Codecs = "opus"

	// CODEC_G722 represents the G.722 codec, offering wideband audio quality.
	CODEC_G722 Codecs = "G722"

	// CODEC_G723 represents the G.723 codec, used for low-bitrate voice compression.
	CODEC_G723 Codecs = "G723"

	// CODEC_ILBC represents the iLBC (Internet Low Bitrate Codec), designed for robust voice transmission over IP.
	CODEC_ILBC Codecs = "iLBC"

	// CODEC_SPEEX represents the Speex codec, an open-source codec optimized for speech.
	CODEC_SPEEX Codecs = "speex"
)

// ICE defines the available modes for handling ICE (Interactive Connectivity Establishment)
// in RTP engine operations. ICE is used to establish peer-to-peer media connections,
// especially in NAT and firewall traversal scenarios.
type ICE string

const (
	// ICERemove disables ICE processing and removes any ICE-related attributes.
	ICERemove ICE = "remove"

	// ICEForce forces ICE negotiation regardless of remote support.
	ICEForce ICE = "force"

	// ICEDefault applies ICE handling based on default engine behavior or configuration.
	ICEDefault ICE = "default"

	// ICEForceRelay forces media to be relayed through ICE candidates (e.g., TURN servers).
	ICEForceRelay ICE = "force-relay"

	// ICEOptional enables ICE negotiation only if supported by the remote party.
	ICEOptional ICE = "optional"
)

// DTLS defines the available modes for handling DTLS (Datagram Transport Layer Security)
// in RTP engine operations. DTLS is used to secure media streams, especially in WebRTC communications.
type DTLS string

const (
	// DTLSOff disables DTLS entirely.
	DTLSOff DTLS = "off"

	// DTLSNo indicates that DTLS is not used for the session.
	DTLSNo DTLS = "no"

	// DTLSDisable disables DTLS negotiation and processing.
	DTLSDisable DTLS = "disable"

	// DTLSPassive sets DTLS to passive mode, waiting for the remote peer to initiate the handshake.
	DTLSPassive DTLS = "passive"

	// DTLSActive sets DTLS to active mode, initiating the handshake with the remote peer.
	DTLSActive DTLS = "active"
)

// DTLSReverse defines the reverse DTLS roles used in RTP engine operations.
// These roles are used to determine how DTLS handshakes are initiated or responded to
// when reversing the media direction or endpoint roles.
type DTLSReverse string

const (
	// DTLSReversePassive sets the reversed DTLS role to passive, waiting for the remote peer to initiate the handshake.
	DTLSReversePassive DTLSReverse = "passive"

	// DTLSReverseActive sets the reversed DTLS role to active, initiating the handshake with the remote peer.
	DTLSReverseActive DTLSReverse = "active"
)

// DTLSFingerprint defines the supported hash algorithms used for DTLS (Datagram Transport Layer Security) fingerprinting.
// DTLS fingerprints are used to verify the identity of peers during secure media session establishment.
type DTLSFingerprint string

const (
	// DTLSFingerprintSha256 uses the SHA-256 hash algorithm for DTLS fingerprinting.
	DTLSFingerprintSha256 DTLSFingerprint = "sha-256"

	// DTLSFingerprintSha1 uses the SHA-1 hash algorithm for DTLS fingerprinting.
	DTLSFingerprintSha1 DTLSFingerprint = "sha-1"

	// DTLSFingerprintSha224 uses the SHA-224 hash algorithm for DTLS fingerprinting.
	DTLSFingerprintSha224 DTLSFingerprint = "sha-224"

	// DTLSFingerprintSha384 uses the SHA-384 hash algorithm for DTLS fingerprinting.
	DTLSFingerprintSha384 DTLSFingerprint = "sha-384"

	// DTLSFingerprintSha512 uses the SHA-512 hash algorithm for DTLS fingerprinting.
	DTLSFingerprintSha512 DTLSFingerprint = "sha-512"
)

// SDES defines the available modes and options for handling SDES (Session Description Protocol Security Descriptions)
// in RTP engine operations. SDES is used to negotiate SRTP (Secure Real-time Transport Protocol) parameters
// through SDP, enabling secure media transmission.
type SDES string

const (
	// SDESOff disables SDES processing.
	SDESOff SDES = "off"

	// SDESNo indicates that SDES is not used for the session.
	SDESNo SDES = "no"

	// SDESDisable disables SDES negotiation and related attributes.
	SDESDisable SDES = "disable"

	// SDESNonew prevents generation of new SDES parameters.
	SDESNonew SDES = "nonew"

	// SDESPad adds padding to SDES crypto attributes.
	SDESPad SDES = "pad"

	// SDESStatic uses static SDES keys instead of dynamically generated ones.
	SDESStatic SDES = "static"

	// SDESPrefer prioritizes SDES over other key exchange mechanisms.
	SDESPrefer SDES = "prefer"

	// SDESUnencrypted_srtp allows unencrypted SRTP streams.
	SDESUnencrypted_srtp SDES = "unencrypted_srtp"

	// SDESUnencrypted_srtcp allows unencrypted SRTCP streams.
	SDESUnencrypted_srtcp SDES = "unencrypted_srtcp"

	// SDESUnauthenticated_srtp allows SRTP without authentication.
	SDESUnauthenticated_srtp SDES = "unauthenticated_srtp"

	// SDESEncrypted_srtp enables encrypted SRTP streams.
	SDESEncrypted_srtp SDES = "encrypted_srtp"

	// SDESEncrypted_srtcp enables encrypted SRTCP streams.
	SDESEncrypted_srtcp SDES = "encrypted_srtcp"

	// SDESAuthenticated_srtp enables authenticated SRTP streams.
	SDESAuthenticated_srtp SDES = "authenticated_srtp"

	// SDESNo* flags disable specific crypto suites from being used in SDES negotiation.
	SDESNoAEAD_AES_256_GCM        SDES = "no-AEAD_AES_256_GCM"
	SDESNoAEAD_AES_128_GCM        SDES = "no-AEAD_AES_128_GCM"
	SDESNoAES_256_CM_HMAC_SHA1_80 SDES = "no-AES_256_CM_HMAC_SHA1_80"
	SDESNoAES_256_CM_HMAC_SHA1_32 SDES = "no-AES_256_CM_HMAC_SHA1_32"
	SDESNoAES_192_CM_HMAC_SHA1_80 SDES = "no-AES_192_CM_HMAC_SHA1_80"
	SDESNoAES_192_CM_HMAC_SHA1_32 SDES = "no-AES_192_CM_HMAC_SHA1_32"
	SDESNoAES_CM_128_HMAC_SHA1_80 SDES = "no-AES_CM_128_HMAC_SHA1_80"
	SDESNoAES_CM_128_HMAC_SHA1_32 SDES = "no-AES_CM_128_HMAC_SHA1_32"
	SDESNoF8_128_HMAC_SHA1_80     SDES = "no-F8_128_HMAC_SHA1_80"
	SDESNoF8_128_HMAC_SHA1_32     SDES = "no-F8_128_HMAC_SHA1_32"
	SDESNoNULL_HMAC_SHA1_80       SDES = "no-NULL_HMAC_SHA1_80"
	SDESNoNULL_HMAC_SHA1_32       SDES = "no-NULL_HMAC_SHA1_32"

	// SDESOnly* flags restrict SDES negotiation to specific crypto suites only.
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

// OSRTP defines the modes for handling Opportunistic SRTP (OSRTP) in RTP engine operations.
// OSRTP allows secure media transmission using SRTP without requiring prior key exchange,
// enabling encryption when both endpoints support it.
type OSRTP string

const (
	// OSRTPOffer initiates an OSRTP negotiation.
	OSRTPOffer OSRTP = "offer"

	// OSRTPOfferRFC initiates OSRTP negotiation using the RFC-compliant method.
	OSRTPOfferRFC OSRTP = "offer-RFC"

	// OSRTPOfferLegacy initiates OSRTP negotiation using legacy signaling methods.
	OSRTPOfferLegacy OSRTP = "offer-legacy"

	// OSRTPAcceptRFC accepts OSRTP using the RFC-compliant method.
	OSRTPAcceptRFC OSRTP = "accept-RFC"

	// OSRTPAcceptLegacy accepts OSRTP using legacy signaling methods.
	OSRTPAcceptLegacy OSRTP = "accept-legacy"

	// OSRTPAccept accepts OSRTP negotiation regardless of the method used.
	OSRTPAccept OSRTP = "accept"
)

// AddressFamily defines the IP address families supported for RTP engine operations.
// These values specify whether IPv4 or IPv6 is used for media and signaling transport.
type AddressFamily string

const (
	// AddressFamilyIP4 represents the IPv4 address family.
	AddressFamilyIP4 AddressFamily = "IP4"

	// AddressFamilyIP6 represents the IPv6 address family.
	AddressFamilyIP6 AddressFamily = "IP6"
)

// Connection defines the types of media connection behaviors used in RTP engine operations.
type Connection string

const (
	// MohConnection represents a "zero" connection, typically used for Music on Hold (MOH) scenarios.
	// This setting may indicate that the media stream should be silenced or replaced with a placeholder,
	// such as music, during call hold states.
	MohConnection Connection = "zero"
)

// Record defines the available options for controlling media recording behavior
// in RTP engine operations. These values determine whether media streams should be recorded.
type Record string

const (
	// RecordYes enables media recording.
	RecordYes Record = "yes"

	// RecordNo disables media recording.
	RecordNo Record = "no"

	// RecordOn turns media recording on (similar to "yes").
	RecordOn Record = "on"

	// RecordOff turns media recording off (similar to "no").
	RecordOff Record = "off"
)
