// Package rtpengine provides functionality for interacting with an RTP engine proxy,
// including connection management, command encoding/decoding, and parameter structures
// for RTP session control.
//
// It defines types for requests and responses, as well as various parameter structs
// for configuring RTP engine operations. The package supports both TCP and UDP connections
// and uses bencode for message serialization.
//
// Main types and functions:
//   - Engine: Manages connection and configuration to the RTP engine proxy.
//   - RequestRtp, ResponseRtp: Structures for command requests and responses.
//   - ParamsOptString, ParamsOptInt, ParamsOptStringArray: Parameter structs for RTP operations.
//   - Conn, ConnUDP: Methods to open TCP/UDP connections to the RTP engine.
//   - EncodeComando: Encodes a command request with a cookie using bencode.
//   - DecodeResposta: Decodes a response from the RTP engine, validating the cookie.
//
// The package relies on external libraries for bencode serialization, UUID generation,
// structured logging, and mapstructure decoding.
package rtpengine

import (
	"bytes"
	"fmt"
	"net"
	"time"

	bencode "github.com/anacrolix/torrent/bencode"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	ben "github.com/stefanovazzocell/bencode"
)

// Engine represents a network engine that manages connections and communication parameters.
// It holds TCP and UDP connections, IP address, port, DNS resolver, protocol type, and a numeric identifier.
type Engine struct {
	con    net.Conn
	conUDP *net.UDPConn
	ip     net.IP
	port   int
	dns    *net.Resolver
	proto  string
	ng     int
}

// RequestRtp represents a request to the RTP engine, containing the command to be executed
// and optional parameters. The struct embeds ParamsOptString, ParamsOptInt, and ParamsOptStringArray
// to allow flexible inclusion of string, integer, and string array options respectively.
// The Command field specifies the action for the RTP engine, and is serialized using both
// JSON and Bencode formats.
type RequestRtp struct {
	Command string `json:"command" bencode:"command"`
	*ParamsOptString
	*ParamsOptInt
	*ParamsOptStringArray
}

// ResponseRtp represents the response structure from the RTP engine.
// It contains information about the result of the RTP operation, including the SDP,
// error and warning messages, timestamps, SSRC, tags, and other metadata related to RTP processing.
// All fields are annotated for both JSON and Bencode serialization.
//
// Fields:
//
//	Result          string      - The result status of the RTP operation.
//	Sdp             string      - The Session Description Protocol (SDP) data.
//	ErrorReason     string      - Error reason message, if any.
//	Warning         string      - Warning message, if any.
//	Created         int         - Timestamp when the response was created.
//	CreatedUs       int         - Microsecond timestamp of creation.
//	LastSignal      int         - Timestamp of the last signal.
//	LastRedisUpdate int         - Timestamp of the last Redis update.
//	SSRC            interface{} - Synchronization source identifier(s).
//	Tags            interface{} - Additional tags or metadata.
//	FromTag         string      - SIP from-tag value.
//	FromTags        []string    - List of SIP from-tag values.
//	ToTag           string      - SIP to-tag value.
//	Totals          TotalRTP    - RTP and RTCP statistics for the session.
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
	FromTag         string      `json:"from-tag,omitempty" bencode:"from-tag,omitempty"`
	FromTags        []string    `json:"from-tags,omitempty" bencode:"from-tags,omitempty"`
	ToTag           string      `json:"to-tag,omitempty" bencode:"to-tag,omitempty"`
	Totals          TotalRTP    `json:"totals,omitempty" bencode:"totals,omitempty"`
}

// TotalRTP represents the total RTP and RTCP statistics for a session.
// It contains two fields, Rtp and Rtcp, each of type ValuesRTP, which hold
// packet, byte, and error counts for RTP and RTCP streams respectively.
// The struct is annotated for both JSON and Bencode serialization.
//
// Fields:
//
//	Rtp  ValuesRTP `json:"RTP,omitempty" bencode:"RTP,omitempty"`  - RTP stream statistics.
//	Rtcp ValuesRTP `json:"RCTP,omitempty" bencode:"RTP,omitempty"` - RTCP stream statistics.
type TotalRTP struct {
	Rtp  ValuesRTP `json:"RTP,omitempty" bencode:"RTP,omitempty"`
	Rtcp ValuesRTP `json:"RCTP,omitempty" bencode:"RTP,omitempty"`
}

// ValuesRTP represents RTP or RTCP stream statistics for a session.
// It contains the number of packets, bytes, and errors for the stream.
// The struct is annotated for both JSON and Bencode serialization.
//
// Fields:
//
//	Packets int - The number of RTP/RTCP packets.
//	Bytes   int - The total number of bytes transmitted.
//	Errors  int - The number of errors encountered.
type ValuesRTP struct {
	Packets int `json:"packets,omitempty" bencode:"packets,omitempty"`
	Bytes   int `json:"bytes,omitempty" bencode:"bytes,omitempty"`
	Errors  int `json:"errors,omitempty" bencode:"errors,omitempty"`
}

// ParamsOptString defines a set of optional parameters for RTP engine operations.
// Each field represents a configurable option that can be serialized to JSON or bencode formats.
// The struct includes tags for both serialization formats and supports various RTP-related settings,
// such as transport protocol, media address, ICE, DTLS, metadata, DTMF, SDP attributes, recording options, and more.
//
// Fields:
//
//	FromTag                string                 - SIP from-tag value.
//	ToTag                  string                 - SIP to-tag value.
//	CallId                 string                 - Call identifier.
//	TransportProtocol      TransportProtocol      - Transport protocol for RTP.
//	MediaAddress           string                 - Media IP address.
//	ICE                    ICE                    - ICE configuration.
//	AddressFamily          AddressFamily          - Address family (e.g., IPv4, IPv6).
//	DTLS                   DTLS                   - DTLS configuration.
//	ViaBranch              string                 - SIP Via branch tag.
//	XmlrpcCallback         string                 - XML-RPC callback URL.
//	Metadata               string                 - Additional metadata.
//	File                   string                 - File path or name.
//	Code                   string                 - Custom code or identifier.
//	DTLSFingerprint        DTLSFingerprint        - DTLS fingerprint.
//	ICELite                string                 - ICE Lite mode.
//	MediaEcho              string                 - Media echo setting.
//	Label                  string                 - Label for identification.
//	SetLabel               string                 - Label to set.
//	FromLabel              string                 - Label for the sender.
//	ToLabel                string                 - Label for the receiver.
//	DTMFSecurity           string                 - DTMF security setting.
//	Digit                  string                 - DTMF digit.
//	DTMFSecurityTrigger    string                 - DTMF security trigger.
//	DTMFSecurityTriggerEnd string                 - DTMF security trigger end.
//	Trigger                string                 - Trigger event.
//	TriggerEnd             string                 - Trigger end event.
//	All                    string                 - Apply to all.
//	Frequency              string                 - Frequency value.
//	Blob                   string                 - Binary large object.
//	Sdp                    string                 - SDP data.
//	AudioPlayer            string                 - Audio player setting.
//	DTMFLogDest            string                 - DTMF log destination.
//	OutputDestination      string                 - Output destination.
//	VscStartRec            string                 - VSC start recording.
//	VscStopRec             string                 - VSC stop recording.
//	VscPauseRec            string                 - VSC pause recording.
//	VscStartStopRec        string                 - VSC start/stop recording.
//	VscPauseResumeRec      string                 - VSC pause/resume recording.
//	VscStartPauseResumeRec string                 - VSC start/pause/resume recording.
//	RtppFlags              string                 - RTP proxy flags.
//	SdpAttr                *ParamsSdpAttrSections - SDP attribute sections.
//	Template               string                 - Template name.
//	RecordCall             Record                 - Call recording option.
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
	Template               string                 `json:"template,omitempty" bencode:"template,omitempty"`
	RecordCall             Record                 `json:"record-call,omitempty" bencode:"record-call,omitempty"`
}

// ParamsOptInt defines a set of integer-based parameters for RTP engine operations.
// Each field represents a configurable option that can be serialized to JSON or bencode formats.
// The struct supports various RTP-related settings, such as TOS, delay, volume, DTMF, ptime, and session duration.
//
// Fields:
//
//	TOS              int - Type of Service (QoS) value.
//	DeleteDelay      int - Delay before deleting the session.
//	DelayBuffer      int - Buffer delay value.
//	Volume           int - Volume adjustment.
//	TriggerEndTime   int - End time for trigger events.
//	TriggerEndDigits int - Number of digits for trigger end.
//	DTMFDelay        int - DTMF signal delay.
//	Ptime            int - Packetization time (ms).
//	PtimeReverse     int - Reverse packetization time (ms).
//	DbId             int - Database identifier.
//	Duration         int - Session duration (seconds).
//	RepeatTimes      int - Repeat times duration (seconds).
//	RepeatDuration   int - Repeat durations (seconds).
//	StartPos         int - Start position (seconds).
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
	RepeatTimes      int `json:"repeat-times,omitempty" bencode:"repeat-times,omitempty"`
	RepeatDuration   int `json:"repeat-duration,omitempty" bencode:"repeat-duration,omitempty"`
	StartPos         int `json:"start-pos,omitempty" bencode:"rstart-pos,omitempty"`
}

// ParamsOptStringArray defines a set of array-based parameters for RTP engine operations.
// Each field represents a configurable option that can be serialized to JSON or bencode formats.
// The struct supports various RTP-related settings, such as flags, RTCP multiplexing, SDES, supported features,
// T.38 fax, OSRTP, received-from addresses, from-tags, frequencies, replace parameters, and music on hold.
//
// Fields:
//
//	Flags        []ParamFlags   - List of flags for RTP operations.
//	RtcpMux      []ParamRTCPMux - RTCP multiplexing options.
//	SDES         []SDES         - SDES encryption options.
//	Supports     []string       - Supported features.
//	T38          []string       - T.38 fax options.
//	OSRTP        []OSRTP        - OSRTP encryption options.
//	ReceivedFrom []string       - List of received-from addresses.
//	FromTags     []string       - List of SIP from-tag values.
//	Frequencies  []string       - List of frequency values.
//	Replace      []ParamReplace - List of parameters to be replaced.
//	Moh          []ParamMoh     - Music On Hold file configurations.
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
	Moh          []ParamMoh     `json:"moh,omitempty" bencode:"moh,omitempty"`
}

// ParamsSdpAttrSections defines the SDP attribute sections for RTP engine operations.
// This struct allows you to specify attribute modifications for different SDP sections,
// including global, audio, video, and none. Each field is a pointer to ParamsSdpAttrCommands,
// which contains lists of attributes to add, remove, or substitute in the respective section.
// Fields are annotated for both JSON and Bencode serialization.
//
// Fields:
//
//	Global *ParamsSdpAttrCommands - Attribute commands for the global SDP section.
//	Audio  *ParamsSdpAttrCommands - Attribute commands for the audio SDP section.
//	Video  *ParamsSdpAttrCommands - Attribute commands for the video SDP section.
//	None   *ParamsSdpAttrCommands - Attribute commands for the none SDP section.
type ParamsSdpAttrSections struct {
	Global *ParamsSdpAttrCommands `json:"global,omitempty" bencode:"global,omitempty"`
	Audio  *ParamsSdpAttrCommands `json:"audio,omitempty" bencode:"audio,omitempty"`
	Video  *ParamsSdpAttrCommands `json:"video,omitempty" bencode:"video,omitempty"`
	None   *ParamsSdpAttrCommands `json:"none,omitempty" bencode:"none,omitempty"`
}

// ParamsSdpAttrCommands defines attribute commands for SDP sections in RTP engine operations.
// This struct allows you to specify lists of SDP attributes to add, remove, or substitute
// within a given SDP section. Fields are annotated for both JSON and Bencode serialization.
//
// Fields:
//
//	Add        []string   - List of SDP attributes to add.
//	Remove     []string   - List of SDP attributes to remove.
//	Substitute [][]string - List of attribute substitutions, where each inner slice contains
//	                       the original attribute and its replacement.
type ParamsSdpAttrCommands struct {
	Add        []string   `json:"add,omitempty" bencode:"add,omitempty"`
	Remove     []string   `json:"remove,omitempty" bencode:"remove,omitempty"`
	Substitute [][]string `json:"substitute,omitempty" bencode:"substitute,omitempty"`
}

// ParamMoh defines the attributes for Music On Hold (MOH) configuration in RTP engine operations.
// This struct allows you to specify the audio file, binary data, database ID, playback mode,
// and connection details for MOH. Fields are annotated for both JSON and Bencode serialization.
//
// Fields:
//
//	File       string     - The path or name of the MOH audio file.
//	Blob       string     - Binary data for the MOH audio.
//	DbId       string     - Database identifier for the MOH resource.
//	Mode       string     - Playback mode (e.g., "sendonly").
//	Connection Connection - Connection details for the MOH resource.
type ParamMoh struct {
	File       string     `json:"file,omitempty" bencode:"file,omitempty"`
	Blob       string     `json:"blob,omitempty" bencode:"blob,omitempty"`
	DbId       string     `json:"db-id,omitempty" bencode:"db-id,omitempty"`
	Mode       string     `json:"mode,omitempty" bencode:"mode,omitempty"`
	Connection Connection `json:"connection,omitempty" bencode:"connection,omitempty"`
}

// GetCookie generates a unique cookie string for command identification.
// This function uses UUID generation to create a random string, which can be used
// to correlate requests and responses when communicating with the RTP engine.
//
// Returns:
//
//	string - A newly generated UUID string to be used as a cookie.
func (r *Engine) GetCookie() string {
	return uuid.NewString()
}

// GetIP returns the default IP address assigned for the RTP engine connection.
// This method retrieves the IP address stored in the Engine instance.
//
// Returns:
//
//	net.IP - The IP address used for the RTP engine connection.
func (r *Engine) GetIP() net.IP {
	return r.ip
}

// GetPort returns the default port assigned for the RTP engine connection.
// This method retrieves the port number stored in the Engine instance.
//
// Returns:
//
//	int - The port number used for the RTP engine connection.
func (r *Engine) GetPort() int {
	return r.port
}

// GetProto returns the default protocol used for the connection.
// This method retrieves the protocol string stored in the Engine instance.
// It can be used to determine the communication protocol currently set.
//
// Returns:
//
//	string - the default protocol for the connection.
func (r *Engine) GetProto() string {
	return r.proto
}

// GetNG returns the default NG port used by the controller.
// This method retrieves the NG port value stored in the Engine instance.
//
// Returns:
//
//	int - the default NG port.
func (r *Engine) GetNG() int {
	return r.ng
}

// Conn establishes a connection with the RTP engine proxy using the configured protocol and port.
// It sets a read timeout of 10 seconds and stores the connection in the Engine instance.
//
// Returns:
//
//	net.Conn - the established network connection.
//	error - an error if the connection fails.
func (r *Engine) Conn() (net.Conn, error) {
	engine := r.ip.String() + ":" + fmt.Sprint(r.port)
	conn, err := net.Dial(r.proto, engine)

	if err != nil {
		log.Debug().Str("Debug ", r.proto+" "+engine).Msg(err.Error())
		return nil, err
	}

	defer net.Dial(r.proto, engine)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	r.con = conn
	return r.con, nil
}

// ConnUDP establishes a UDP connection with the RTP engine proxy using the configured protocol and port.
// It sets a read timeout of 10 seconds and stores the UDP connection in the Engine instance.
//
// Returns:
//
//	*net.UDPConn - the established UDP connection.
//	error - an error if the connection fails.
func (r *Engine) ConnUDP() (*net.UDPConn, error) {
	engine := r.ip.String() + ":" + fmt.Sprint(r.port)
	addr := &net.UDPAddr{
		IP:   r.ip,
		Port: r.port,
	}
	conn, err := net.DialUDP(r.proto, nil, addr)

	if err != nil {
		log.Debug().Str("Debug ", r.proto+" "+engine).Msg(err.Error())
		return nil, err
	}

	defer net.DialUDP(r.proto, nil, addr)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	r.conUDP = conn
	return r.conUDP, nil
}

// EncodeComando encodes a command into bencode format and prepends the cookie.
// This function marshals the RequestRtp struct into bencode and combines it with the cookie.
//
// Parameters:
//
//	cookie - a string used for identifying the request.
//	command - a pointer to the RequestRtp struct containing the command data.
//
// Returns:
//
//	[]byte - the encoded command with the cookie.
//	error - an error if encoding fails.
func EncodeComando(cookie string, command *RequestRtp) ([]byte, error) {
	data, err := bencode.Marshal(command)
	if err != nil {
		return nil, err
	}

	bind := []byte(cookie + " ")
	return append(bind, data...), nil
}

// DecodeResposta decodes the response from the RTP engine and validates the cookie.
// It parses the bencoded response and maps it to the ResponseRtp struct.
//
// Parameters:
//
//	cookie - the expected cookie string.
//	resposta - the raw response bytes from the RTP engine.
//
// Returns:
//
//	*ResponseRtp - the decoded response with result and error information.
func DecodeResposta(cookie string, resposta []byte) *ResponseRtp {
	resp := &ResponseRtp{}
	cookieIndex := bytes.IndexAny(resposta, " ")
	if cookieIndex != len(cookie) {
		resp.Result = "error"
		resp.ErrorReason = "Failed to parse the message"
		return resp
	}

	cookieResponse := string(resposta[:cookieIndex])
	if cookieResponse != cookie {
		resp.Result = "error"
		resp.ErrorReason = "Cookie mismatch"
		return resp
	}

	encodedData := string(resposta[cookieIndex+1:])
	decodedDataRaw, err := ben.NewParserFromString(encodedData).AsDict()

	if err != nil {
		return resp
	}

	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &resp,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(decodedDataRaw)
	return resp
}
