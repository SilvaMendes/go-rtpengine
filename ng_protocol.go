package rtpengine

import "fmt"

// ParametrosOption defines a function type used to customize a RequestRtp instance.
// Each ParametrosOption receives a pointer to a RequestRtp and applies specific configuration,
// such as setting flags, codecs, transport protocol, or other SDP-related parameters.
// If the option encounters an error during configuration, it should return the error.
//
// This pattern allows flexible and composable configuration of RTP requests,
// enabling multiple options to be applied in sequence.
//
// Example usage:
//   req, err := SDPOffering(params, req.SetCodecMask([]Codecs{"PCMU"}), req.SetTransportProtocol(RTP_AVP))
type ParametrosOption func(c *RequestRtp) error

// SDPOffering creates an SDP offer request for rtpengine.
// This function initializes a RequestRtp structure with the "Offer" command and the provided parameters.
// It applies any additional options (ParametrosOption) to customize the request, such as codec masks, flags, or transport protocol.
// If any option returns an error, the function aborts and returns the error.
//
// Parameters:
//   parametros *ParamsOptString - The main SDP parameters for the offer.
//   options ...ParametrosOption - Optional functions to further configure the RequestRtp.
//
// Returns:
//   *RequestRtp - The fully configured RTP request for the offer.
//   error       - Any error encountered while applying the options.
//
// Example usage:
//   req, err := SDPOffering(params, req.SetCodecMask([]Codecs{"PCMU"}), req.SetTransportProtocol(RTP_AVP))
func SDPOffering(parametros *ParamsOptString, options ...ParametrosOption) (*RequestRtp, error) {
	request := &RequestRtp{
		Command:              fmt.Sprint(Offer),
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	for _, o := range options {
		if err := o(request); err != nil {
			return nil, err
		}
	}
	return request, nil
}

// SDPAnswer creates an SDP answer request for rtpengine.
// This function initializes a RequestRtp structure with the "Answer" command and the provided parameters.
// It applies any additional options (ParametrosOption) to customize the request, such as codec masks, flags, or transport protocol.
// If any option returns an error, the function aborts and returns the error.
//
// Parameters:
//   parametros *ParamsOptString - The main SDP parameters for the answer.
//   options ...ParametrosOption - Optional functions to further configure the RequestRtp.
//
// Returns:
//   *RequestRtp - The fully configured RTP request for the answer.
//   error       - Any error encountered while applying the options.
//
// Example usage:
//   req, err := SDPAnswer(params, req.SetCodecMask([]Codecs{"PCMU"}), req.SetTransportProtocol(RTP_AVP))
func SDPAnswer(parametros *ParamsOptString, options ...ParametrosOption) (*RequestRtp, error) {
	request := &RequestRtp{
		Command:              fmt.Sprint(Answer),
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	for _, o := range options {
		if err := o(request); err != nil {
			return nil, err
		}
	}
	return request, nil
}

// SDPDelete creates an SDP delete request for rtpengine.
// This function initializes a RequestRtp structure with the "Delete" command and the provided parameters.
// It applies any additional options (ParametrosOption) to customize the request, such as flags or protocol settings.
// If any option returns an error, the function aborts and returns the error.
//
// Parameters:
//   parametros *ParamsOptString - The main SDP parameters for the delete operation.
//   options ...ParametrosOption - Optional functions to further configure the RequestRtp.
//
// Returns:
//   *RequestRtp - The fully configured RTP request for the delete operation.
//   error       - Any error encountered while applying the options.
//
// Example usage:
//   req, err := SDPDelete(params, req.SetFlags([]ParamFlags{"flag1"}))
func SDPDelete(parametros *ParamsOptString, options ...ParametrosOption) (*RequestRtp, error) {
	request := &RequestRtp{
		Command:              fmt.Sprint(Delete),
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	for _, o := range options {
		if err := o(request); err != nil {
			return nil, err
		}
	}
	return request, nil
}

// SetFlags adds a list of flags to the RTP request.
// This function receives a slice of ParamFlags and appends them to the Flags field
// of ParamsOptStringArray in the RequestRtp structure. These flags are used to
// customize the behavior of the RTP engine during SDP negotiation.
//
// Parameters:
//   flags []ParamFlags - List of flags to be added.
//
// Returns:
//   ParametrosOption - A function that applies the flags to the RequestRtp structure.
//
// Example usage:
//   req.SetFlags([]ParamFlags{"flag1", "flag2"})
func (c *RequestRtp) SetFlags(flags []ParamFlags) ParametrosOption {
	return func(s *RequestRtp) error {
		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, flags...)
		return nil
	}
}

// SetTransportProtocol sets the transport protocol for the RTP request.
// This function receives a TransportProtocol value and assigns it to the TransportProtocol field
// of the RequestRtp structure. The transport protocol determines how RTP packets are transmitted
// (e.g., RTP/AVP, RTP/SAVP, UDP/TLS/RTP/SAVP).
//
// Parameters:
//   proto TransportProtocol - The transport protocol to be set.
//
// Returns:
//   ParametrosOption - A function that applies the transport protocol to the RequestRtp structure.
//
// Example usage:
//   req.SetTransportProtocol("RTP/AVP")
func (c *RequestRtp) SetTransportProtocol(proto TransportProtocol) ParametrosOption {
	return func(s *RequestRtp) error {
		s.TransportProtocol = proto
		return nil
	}
}

// SetReplace sets the list of SDP parameters to be replaced in the RTP request.
// This function receives a slice of ParamReplace and assigns it to the Replace field
// of the RequestRtp structure. These parameters specify which SDP attributes should be replaced
// during the negotiation process.
//
// Parameters:
//   replace []ParamReplace - List of SDP parameters to be replaced.
//
// Returns:
//   ParametrosOption - A function that applies the replace parameters to the RequestRtp structure.
//
// Example usage:
//   req.SetReplace([]ParamReplace{"origin", "session"})
func (c *RequestRtp) SetReplace(replace []ParamReplace) ParametrosOption {
	return func(s *RequestRtp) error {
		s.Replace = replace
		return nil
	}
}

// SetRtcpMux sets the RTCP multiplexing options for the RTP request.
// This function receives a slice of ParamRTCPMux and assigns it to the RtcpMux field
// of the RequestRtp structure. These options control how RTCP packets are multiplexed with RTP.
//
// Parameters:
//   rtcpmux []ParamRTCPMux - List of RTCP multiplexing options to be set.
//
// Returns:
//   ParametrosOption - A function that applies the RTCP multiplexing options to the RequestRtp structure.
//
// Example usage:
//   req.SetRtcpMux([]ParamRTCPMux{"demux"})
func (c *RequestRtp) SetRtcpMux(rtcpmux []ParamRTCPMux) ParametrosOption {
	return func(s *RequestRtp) error {
		s.RtcpMux = rtcpmux
		return nil
	}
}

// SetCodecEncoder adds codec transcoding flags to the RTP request.
// This function receives a list of codecs and generates flags in the format "codec-transcode-<codec>",
// which are appended to the Flags field of ParamsOptStringArray.
// These flags instruct rtpengine to transcode the specified codecs during SDP negotiation.
//
// Parameters:
//   codecs []Codecs - List of codecs to be transcoded.
//
// Returns:
//   ParametrosOption - A function that applies the codec transcoding flags to the RequestRtp structure.
//
// Example usage:
//   req.SetCodecEncoder([]Codecs{"PCMU", "PCMA"})
func (c *RequestRtp) SetCodecEncoder(codecs []Codecs) ParametrosOption {
	return func(s *RequestRtp) error {
		trascoder := make([]ParamFlags, 0)
		for _, o := range codecs {
			trascoder = append(trascoder, ParamFlags("codec-transcode-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, trascoder...)
		return nil
	}
}

// SetCodecMask adds codec mask flags to the RTP request.
// This function receives a list of codecs and generates flags in the format "codec-mask-<codec>",
// which are appended to the Flags field of ParamsOptStringArray.
// These flags instruct rtpengine to accept only the specified codecs during SDP negotiation.
//
// Parameters:
//   codecs []Codecs - List of codecs to be masked.
//
// Returns:
//   ParametrosOption - A function that applies the codec mask flags to the RequestRtp structure.
//
// Example usage:
//   req.SetCodecMask([]Codecs{"PCMU", "PCMA"})
func (c *RequestRtp) SetCodecMask(codecs []Codecs) ParametrosOption {
	return func(s *RequestRtp) error {
		mask := make([]ParamFlags, 0)
		for _, o := range codecs {
			mask = append(mask, ParamFlags("codec-mask-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, mask...)
		return nil
	}
}

// SetCodecStrip adds codec strip flags to the RTP request.
// This function receives a list of codecs and generates flags in the format "codec-strip-<codec>",
// which are appended to the Flags field of ParamsOptStringArray.
// These flags instruct rtpengine to remove the specified codecs from the SDP negotiation.
//
// Parameters:
//   codecs []Codecs - List of codecs to be stripped.
//
// Returns:
//   ParametrosOption - A function that applies the codec strip flags to the RequestRtp structure.
//
// Example usage:
//   req.SetCodecStrip([]Codecs{"PCMU", "PCMA"})
func (c *RequestRtp) SetCodecStrip(codecs []Codecs) ParametrosOption {
	return func(s *RequestRtp) error {
		strip := make([]ParamFlags, 0)
		for _, o := range codecs {
			strip = append(strip, ParamFlags("codec-strip-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, strip...)
		return nil
	}
}

// SetCodecExcept adds codec exception flags to the RTP request.
// This function receives a list of codecs and generates flags in the format "codec-except-<codec>",
// which are appended to the Flags field of ParamsOptStringArray.
// These flags instruct rtpengine to exclude the specified codecs from the SDP negotiation.
//
// Parameters:
//   codecs []Codecs - List of codecs to be excluded.
//
// Returns:
//   ParametrosOption - A function that applies the codec exception flags to the RequestRtp structure.
//
// Example usage:
//   req.SetCodecExcept([]Codecs{"PCMU", "PCMA"})
func (c *RequestRtp) SetCodecExcept(codecs []Codecs) ParametrosOption {
	return func(s *RequestRtp) error {
		except := make([]ParamFlags, 0)
		for _, o := range codecs {
			except = append(except, ParamFlags("codec-except-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, except...)
		return nil
	}
}

// DisablesSDES disables SDES (Session Description Protocol Security Descriptions) for the RTP request.
// This function appends the SDESOff flag to the SDES field of ParamsOptStringArray,
// instructing rtpengine to turn off SDES encryption for the session.
//
// Returns:
//   ParametrosOption - A function that applies the SDES disable flag to the RequestRtp structure.
//
// Example usage:
//   req.DisablesSDES()
func (c *RequestRtp) DisablesSDES() ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]SDES, 0)
		sdes = append(sdes, SDESOff)
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// DeleteSDES disables specific SDES crypto suites for the RTP request.
// This function receives a list of CryptoSuite values and generates flags in the format "no-<crypto>",
// which are appended to the SDES field of ParamsOptStringArray.
// These flags instruct rtpengine to disable the specified SDES encryption algorithms for the session.
//
// Parameters:
//   cript []CryptoSuite - List of crypto suites to be disabled.
//
// Returns:
//   ParametrosOption - A function that applies the SDES disable flags to the RequestRtp structure.
//
// Example usage:
//   req.DeleteSDES([]CryptoSuite{"AES_CM_128_HMAC_SHA1_80"})
func (c *RequestRtp) DeleteSDES(cript []CryptoSuite) ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]SDES, 0)
		for _, o := range cript {
			sdes = append(sdes, "no-"+SDES(o))
		}
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// EnableSDES enables specific SDES crypto suites for the RTP request.
// This function receives a list of CryptoSuite values and generates flags in the format "only-<crypto>",
// which are appended to the SDES field of ParamsOptStringArray.
// These flags instruct rtpengine to use only the specified SDES encryption algorithms for the session.
//
// Parameters:
//   cript []CryptoSuite - List of crypto suites to be enabled.
//
// Returns:
//   ParametrosOption - A function that applies the SDES enable flags to the RequestRtp structure.
//
// Example usage:
//   req.EnableSDES([]CryptoSuite{"AES_CM_128_HMAC_SHA1_80"})
func (c *RequestRtp) EnableSDES(cript []CryptoSuite) ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]SDES, 0)
		for _, o := range cript {
			sdes = append(sdes, "only-"+SDES(o))
		}
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// ICERemove removes all ICE attributes from the SDP in the RTP request.
// This function sets the ICE field of the RequestRtp structure to "remove",
// indicating that all ICE candidates and related attributes should be stripped from the SDP
// before sending it to rtpengine. This ensures that the media proxy becomes the sole ICE candidate.
//
// Returns:
//   ParametrosOption - A function that applies the ICE removal to the RequestRtp structure.
//
// Example usage:
//   req.ICERemove()
func (c *RequestRtp) ICERemove() ParametrosOption {
	return func(s *RequestRtp) error {
		s.ICE = "remove"
		return nil
	}
}

// ICEForce forces the use of ICE attributes in the SDP for the RTP request.
// This function sets the ICE field of the RequestRtp structure to "force",
// instructing rtpengine to enforce ICE negotiation regardless of the original SDP content.
//
// Returns:
//   ParametrosOption - A function that applies the ICE force option to the RequestRtp structure.
//
// Example usage:
//   req.ICEForce()
func (c *RequestRtp) ICEForce() ParametrosOption {
	return func(s *RequestRtp) error {
		s.ICE = "force"
		return nil
	}
}

// SetAttrChange sets the SDP attribute sections for the RTP request.
// This function receives a pointer to ParamsSdpAttrSections and assigns it to the SdpAttr field
// of the RequestRtp structure. It allows you to specify which SDP attributes should be added, removed,
// or replaced during the negotiation process.
//
// Parameters:
//   sdpAttr *ParamsSdpAttrSections - Pointer to the SDP attribute sections to be applied.
//
// Returns:
//   ParametrosOption - A function that applies the SDP attribute changes to the RequestRtp structure.
//
// Example usage:
//   req.SetAttrChange(&ParamsSdpAttrSections{/* attribute modifications */})
func (c *RequestRtp) SetAttrChange(sdpAttr *ParamsSdpAttrSections) ParametrosOption {
	return func(s *RequestRtp) error {
		s.SdpAttr = sdpAttr
		return nil
	}
}

// SetViaBranchTag sets the Via branch tag for the RTP request.
// This function assigns the provided branch string to the ViaBranch field
// of the RequestRtp structure. The Via branch tag is used to uniquely identify
// SIP transactions and is important for proper routing and correlation of SIP messages.
//
// Parameters:
//   branch string - The branch tag to be set.
//
// Returns:
//   ParametrosOption - A function that applies the Via branch tag to the RequestRtp structure.
//
// Example usage:
//   req.SetViaBranchTag("z9hG4bK-branchvalue")
func (c *RequestRtp) SetViaBranchTag(branch string) ParametrosOption {
	return func(s *RequestRtp) error {
		s.ViaBranch = branch
		return nil
	}
}

// SetPtimeCodecOffer sets the ptime value for the codec in the SDP offer.
// This function assigns the provided integer value to the Ptime field of the RequestRtp structure.
// The ptime parameter defines the packetization time (in milliseconds) for the codec in the offer,
// which can be used to control the duration of audio frames in RTP packets.
//
// Parameters:
//   ptime int - The packetization time value to be set.
//
// Returns:
//   ParametrosOption - A function that applies the ptime value to the RequestRtp structure.
//
// Example usage:
//   req.SetPtimeCodecOffer(20)
func (c *RequestRtp) SetPtimeCodecOffer(ptime int) ParametrosOption {
	return func(s *RequestRtp) error {
		s.Ptime = ptime
		return nil
	}
}

// SetPtimeCodecAnswer sets the ptime value for the codec in the SDP answer.
// This function assigns the provided integer value to the PtimeReverse field of the RequestRtp structure.
// The ptime parameter defines the packetization time (in milliseconds) for the codec in the answer,
// which can be used to control the duration of audio frames in RTP packets.
//
// Parameters:
//   ptime int - The packetization time value to be set.
//
// Returns:
//   ParametrosOption - A function that applies the ptime value to the RequestRtp structure.
//
// Example usage:
//   req.SetPtimeCodecAnswer(20)
func (c *RequestRtp) SetPtimeCodecAnswer(ptime int) ParametrosOption {
	return func(s *RequestRtp) error {
		s.PtimeReverse = ptime
		return nil
	}
}

// SetReceivedFrom adds the received-from attribute to the RTP request.
// This function is used when SDP addresses are not reliable. It receives the address family and address,
// and appends them to the ReceivedFrom field of the RequestRtp structure.
//
// Parameters:
//   addressFamily AddressFamily - The address family (e.g., "ipv4", "ipv6").
//   Address string - The IP address to be set.
//
// Returns:
//   ParametrosOption - A function that applies the received-from attribute to the RequestRtp structure.
//
// Example usage:
//   req.SetReceivedFrom(ipv4, "192.168.1.100")
func (c *RequestRtp) SetReceivedFrom(addressFamily AddressFamily, Address string) ParametrosOption {
	return func(s *RequestRtp) error {
		receivedFrom := make([]string, 0)
		s.ReceivedFrom = append(receivedFrom, string(addressFamily), Address)
		return nil
	}
}

// SetMediaAddress sets the media address for the RTP request.
// This function assigns the provided IP address string to the MediaAddress field
// of the RequestRtp structure. The media address is used to specify the destination
// IP address for media streams in the SDP negotiation.
//
// Parameters:
//   Address string - The IP address to be set as the media address.
//
// Returns:
//   ParametrosOption - A function that applies the media address to the RequestRtp structure.
//
// Example usage:
//   req.SetMediaAddress("192.168.1.200")
func (c *RequestRtp) SetMediaAddress(Address string) ParametrosOption {
	return func(s *RequestRtp) error {
		s.MediaAddress = Address
		return nil
	}
}

// RecordYes enables call recording for the RTP request.
// This function sets the RecordCall field of the RequestRtp structure to "yes",
// instructing rtpengine to record the media stream for the session.
//
// Returns:
//   ParametrosOption - A function that applies the call recording option to the RequestRtp structure.
//
// Example usage:
//   req.RecordYes()
func (c *RequestRtp) RecordYes() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "yes"
		return nil
	}
}

// RecordNo disables call recording for the RTP request.
// This function sets the RecordCall field of the RequestRtp structure to "no",
// instructing rtpengine not to record the media stream for the session.
//
// Returns:
//   ParametrosOption - A function that applies the call recording disable option to the RequestRtp structure.
//
// Example usage:
//   req.RecordNo()
func (c *RequestRtp) RecordNo() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "no"
		return nil
	}
}

// RecordOn enables call recording for the RTP request.
// This function sets the RecordCall field of the RequestRtp structure to "on",
// instructing rtpengine to start recording the media stream for the session.
//
// Returns:
//   ParametrosOption - A function that applies the call recording option to the RequestRtp structure.
//
// Example usage:
//   req.RecordOn()
func (c *RequestRtp) RecordOn() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "on"
		return nil
	}
}

// RecordOff disables call recording for the RTP request.
// This function sets the RecordCall field of the RequestRtp structure to "off",
// instructing rtpengine to stop recording the media stream for the session.
//
// Returns:
//   ParametrosOption - A function that applies the call recording off option to the RequestRtp structure.
//
// Example usage:
//   req.RecordOff()
func (c *RequestRtp) RecordOff() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "off"
		return nil
	}
}

// SetMohFile adds a Music On Hold (MOH) file to the RTP request.
// This function appends a ParamMoh struct with the specified file and "sendonly" mode
// to the Moh field of the RequestRtp structure. It is used to configure the RTP engine
// to play a specific audio file as music on hold during the session.
//
// Parameters:
//   file string - The path or name of the MOH audio file to be played.
//
// Returns:
//   ParametrosOption - A function that applies the MOH file configuration to the RequestRtp structure.
//
// Example usage:
//   req.SetMohFile("holdmusic.wav")
func (c *RequestRtp) SetMohFile(file string) ParametrosOption {
	return func(s *RequestRtp) error {
		s.Moh = append(s.Moh, ParamMoh{File: file, Mode: "sendonly"})
		return nil
	}
}
