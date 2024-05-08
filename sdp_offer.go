package rtpengine

type Menssagem string

const (
	SDP                    Menssagem = "sdp"
	CallId                 Menssagem = "call-id"
	FromTag                Menssagem = "from-tag"
	All                    Menssagem = "all"
	None                   Menssagem = "none"
	OfferAnswer            Menssagem = "offer-answer"
	ExceptOfferAnswer      Menssagem = "except-offer-answer"
	Flows                  Menssagem = "flows"
	AddressFamily          Menssagem = "address family"
	AudioPlayer            Menssagem = "audio player"
	DelayBuffer            Menssagem = "delay-buffer"
	Direction              Menssagem = "direction"
	Digit                  Menssagem = "digit"
	DropTraffic            Menssagem = "drop-traffic"
	DTLS                   Menssagem = "DTLS"
	DTLSReverse            Menssagem = "DTLS-reverse"
	DTLSFingerprint        Menssagem = "DTLS-fingerprint"
	DTMFsecurity           Menssagem = "DTMF-security"
	DTMFSecurityTrigge     Menssagem = "DTMF-security-trigge"
	DTMFSecurityTriggerEnd Menssagem = "DTMF-security-trigger-end"
	DTMFDelay              Menssagem = "DTMF-delay"
	DTMFLogDest            Menssagem = "DTMF-log-dest"
	Frequency              Menssagem = "frequency"
	EndpointLearning       Menssagem = "endpoint-learning"
	FromTags               Menssagem = "from-tags"
	FromTagsFlags          Menssagem = "from-tags-"
	GenerateRTCP           Menssagem = "generate RTCP"
	ICE                    Menssagem = "ICE"
	ICELite                Menssagem = "ICE-lite"
	Interface              Menssagem = "interface"
	FromLabel              Menssagem = "from-label"
	Label                  Menssagem = "label"
)

// O dicionário de solicitação
type MsgOffer struct {
	Sdp            string
	Callid         string
	Fromtag        string
	OfferOprions   MsgOfferOptions
	Seguranca      MsgDTLS
	DTMF           MsgDTMF
	EndpointsFlags MsgEndpointsFlags
}

// chaves opcional
type MsgOfferOptions struct {
	Opt           string // optional {all,none,offer-answer,except-offer-answer,flows}
	AddressFamily string // optional {IP4,IP6}
	AudioPlayer   string // optional {default,transcoding,off,always}
	DelayBuffer   int
	Direction     string // optional {e,i,round-robin-calls, string definida no parametro --interface}
	Digit         string // optional Define o dígito de substituição para DTMF-security=DTMF.
	DropTraffic   string // optional {start,stop}
}

// Influencia o comportamento do DTLS-SRTP
type MsgDTLS struct {
	DTLS            string // optional {off,no,disable,passive,active}
	DTLSReverse     string // optional {passive,active}
	DTLSFingerprint string // optional {SHA-1,SHA-224, SHA-256, SHA-384, SHA-512}
}

// Influencia o comportamento do DTMF
type MsgDTMF struct {
	DTMFsecurity           string // optional {block} ou DTMF drop,silence,tone,random,zero random DTMF zero,off
	DTMFSecurityTrigge     string // optional {trigger}
	DTMFSecurityTriggerEnd string // optional {end trigger}
	DTMFDelay              int    // optional int msec delay
	DTMFLogDest            string // optional endereço para capturao o envio do DTMF para sessão formato host:port
	Frequency              string // optional string "400,450,800" lista de frequencia hz
}

// Manipular comportamento do ednpoints
type MsgEndpointsFlags struct {
	EndpointLearning string // optional{off,immediate,delayed,heuristic}
	FromTags         string // optional Contém uma lista de strings usadas para selecionar vários participantes cada uma prefixada com subscribe request flags from-tags-
	GenerateRTCP     string // optional {on,off}
	ICE              string // optional {remove,force,default,force-relay,optional}
	ICELite          string // optional {forward,backward,both,off,offer}
	Interface        string // optional  string com o nome de uma das interfaces, assim como "direction"
	FromLabel        string //optional so tem efeito para log
	Label            string //optional so tem efeito para log identioc ao FromLabel
}
