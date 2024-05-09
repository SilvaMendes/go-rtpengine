package rtpengine

// Parametros usado como passagem de flags
type ParamsFlags struct {
	TrustAddress         string `json:"trust-address,omitempty"`
	Symmetric            string `json:"symmetric,omitempty"`
	Asymmetric           string `json:"asymmetric,omitempty"`
	Unidirectional       string `json:"unidirectional,omitempty"`
	Force                string `json:"force,omitempty"`
	StrictSource         string `json:"strict-source,omitempty"`
	MediaHandover        string `json:"media-handover,omitempty"`
	SipSourceAddress     string `json:"sip-source-address,omitempty"`
	Reset                string `json:"reset,omitempty"`
	PortLatching         string `json:"port-latching,omitempty"`
	NoRtcpAttribute      string `json:"no-rtcp-attribute,omitempty"`
	FullRtcpAttribute    string `json:"full-rtcp-attribute,omitempty"`
	LoopProtect          string `json:"loop-protect,omitempty"`
	RecordCall           string `json:"record-call,omitempty"`
	AlwaysTranscode      string `json:"always-transcode,omitempty"`
	SIPREC               string `json:"SIPREC,omitempty"`
	PadCrypto            string `json:"pad-crypto,omitempty"`
	GenerateMid          string `json:"generate-mid,omitempty"`
	Fragment             string `json:"fragment,omitempty"`
	OriginalSendrecv     string `json:"original-sendrecv,omitempty"`
	SymmetricCodecs      string `json:"symmetric-codecs,omitempty"`
	AsymmetricCodecs     string `json:"asymmetric-codecs,omitempty"`
	InjectDTMF           string `json:"inject-DTMF,omitempty"`
	DetectDTMF           string `json:"detect-DTMF,omitempty"`
	GenerateRTCP         string `json:"generate-RTCP,omitempty"`
	SingleCodec          string `json:"single-codec,omitempty"`
	NoCodecRenegotiation string `json:"no-codec-renegotiation,omitempty"`
	PierceNAT            string `json:"pierce-NAT,omitempty"`
	SIPSourceAddress     string `json:"SIP-source-address,omitempty"`
	AllowTranscoding     string `json:"allow-transcoding,omitempty"`
	TrickleICE           string `json:"trickle-ICE,omitempty"`
	RejectICE            string `json:"reject-ICE,omitempty"`
	Egress               string `json:"egress,omitempty"`
	NoJitterBuffer       string `json:"no-jitter-buffer,omitempty"`
	Passthrough          string `json:"passthrough,omitempty"`
	NoPassthrough        string `json:"no-passthrough,omitempty"`
	Pause                string `json:"pause,omitempty"`
	EarlyMedia           string `json:"early-media,omitempty"`
	BlockShort           string `json:"block-short,omitempty"`
	RecordingVsc         string `json:"recording-vsc,omitempty"`
	BlockEgress          string `json:"block-egress,omitempty"`
}

// Parametros de comportamento
type ParamsOptString struct {
	FromTag                string `json:"from-tag"`
	ToTag                  string `json:"to-tag"`
	CallId                 string `json:"call-id"`
	TransportProtocol      string `json:"transport-protocol"`
	MediaAddress           string `json:"media-address,omitempty"`
	ICE                    string `json:"ICE,omitempty"`
	AddressFamily          string `json:"address-family,omitempty"`
	DTLS                   string `json:"DTLS,omitempty"`
	ViaBranch              string `json:"via-branch,omitempty"`
	XmlrpcCallback         string `json:"xmlrpc-callback,omitempty"`
	Metadata               string `json:"metadata,omitempty"`
	Address                string `json:"address,omitempty"`
	File                   string `json:"file,omitempty"`
	Code                   string `json:"code,omitempty"`
	DTLSFingerprint        string `json:"DTLS-fingerprint,omitempty"`
	ICELite                string `json:"ICE-lite,omitempty"`
	MediaEcho              string `json:"media-echo,omitempty"`
	Label                  string `json:"label,omitempty"`
	SetLabel               string `json:"set-label,omitempty"`
	FromLabel              string `json:"from-label,omitempty"`
	ToLabel                string `json:"to-label,omitempty"`
	DTMFSecurity           string `json:"DTMF-security,omitempty"`
	Digit                  string `json:"digit,omitempty"`
	DTMFSecurityTrigger    string `json:"DTMF-security-trigger,omitempty"`
	DTMFSecurityTriggerEnd string `json:"DTMF-security-trigger-end,omitempty"`
	Trigger                string `json:"trigger,omitempty"`
	TriggerEnd             string `json:"trigger-end,omitempty"`
	All                    string `json:"all,omitempty"`
	Frequency              string `json:"frequency,omitempty"`
	Blob                   string `json:"blob,omitempty"`
	Sdp                    string `json:"sdp"`
	AudioPlayer            string `json:"audio-player,omitempty"`
	DTMFLogDest            string `json:"dtmf-log-dest,omitempty"`
	OutputDestination      string `json:"output-destination,omitempty"`
	VscStartRec            string `json:"vsc-start-rec,omitempty"`
	VscStopRec             string `json:"vsc-stop-rec,omitempty"`
	VscPauseRec            string `json:"vsc-pause-rec,omitempty"`
	VscStartStopRec        string `json:"vsc-start-stop-rec,omitempty"`
	VscPauseResumeRec      string `json:"vsc-pause-resume-rec,omitempty"`
	VscStartPauseResumeRec string `json:"vsc-start-pause-resume-rec,omitempty"`
	RtppFlags              string `json:"rtpp-flags,omitempty"`
}

// Parametros de comportamento tipo inteiro
type ParamsOptInt struct {
	TOS              int `json:"TOS,omitempty"`
	DeleteDelay      int `json:"delete-delay,omitempty"`
	DelayBuffer      int `json:"delay-buffer,omitempty"`
	Volume           int `json:"volume,omitempty"`
	TriggerEndTime   int `json:"trigger-end-time,omitempty"`
	TriggerEndDigits int `json:"trigger-end-digits,omitempty"`
	DTMFDelay        int `json:"DTMF-delay,omitempty"`
	Ptime            int `json:"ptime,omitempty"`
	DbId             int `json:"db-id,omitempty"`
	Duration         int `json:"duration,omitempty"`
}

// Parametros de comportamento tipo array separado por ','
type ParamsOptStringArray struct {
	Flags        string `json:"flags,omitempty"`
	RtcpMux      string `json:"rtcp-mux,omitempty"`
	SDES         string `json:"SDES,omitempty"`
	Supports     string `json:"supports,omitempty"`
	T38          string `json:"T38,omitempty"`
	OSRTP        string `json:"OSRTP,omitempty"`
	ReceivedFrom string `json:"received-from,omitempty"`
	FromTags     string `json:"from-tags,omitempty"`
	Frequencies  string `json:"frequencies,omitempty"`
}

// Parametros de manipulação dos codecs na oferta
type ParamsOptCodec struct {
	Strip     string `json:"strip,omitempty"`
	Offer     string `json:"offer,omitempty"`
	Transcode string `json:"transcode,omitempty"`
	Mask      string `json:"mask,omitempty"`
	Set       string `json:"set,omitempty"`
	Consume   string `json:"consume,omitempty"`
	Accept    string `json:"accept,omitempty"`
	Except    string `json:"except,omitempty"`
}

// Parametros de substituição de valores no corpo do SDP
type ParamsOptReplace struct {
	Origin                   string `json:"origin,omitempty"`
	SessionConnection        string `json:"session-connection,omitempty"`
	SdpVersion               string `json:"sdp-version,omitempty"`
	Username                 string `json:"username,omitempty"`
	SessionName              string `json:"session-name,omitempty"`
	ZeroAddress              string `json:"zero-address,omitempty"`
	ForceIncrementSdpVersion string `json:"force-increment-sdp-version,omitempty"`
}

// Parametros de manipulação de sessão
type ParamsSdpAttrSections struct {
	Global string `json:"global,omitempty"`
	Audio  string `json:"audio,omitempty"`
	Video  string `json:"video,omitempty"`
}

// Parametros de atributos de comandos
type ParamsSdpAttrCommands struct {
	Add    string `json:"add,omitempty"`
	Remove string `json:"remove,omitempty"`
}
