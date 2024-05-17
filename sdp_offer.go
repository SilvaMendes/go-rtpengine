package rtpengine

// Perfil para o protocolo UDP
func ProfilerRTP_UDP_Offer(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como RTP/AVP
	parametros.TransportProtocol = RTP_AVP

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]ParamFlags, 0)
	sdes := make([]SDES, 0)

	rtcpmux = append(rtcpmux, RTCPDemux)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, StripExtmap, NoRtcpAttribute)
	sdes = append(sdes, SDESOff)

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = ICERemove
	request.DTLS = DTLSOff
	request.SDES = sdes

	return request
}

// Perfil para o protocolo TCP
func ProfilerRTP_TCP_Offer(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como RTP/AVP
	parametros.TransportProtocol = RTP_AVP

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]ParamFlags, 0)
	osrtp := make([]OSRTP, 0)

	rtcpmux = append(rtcpmux, RTCPDemux)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, LoopProtect, StrictSource)
	osrtp = append(osrtp, OSRTPOffer)

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = ICERemove
	request.DTLS = DTLSOff
	request.OSRTP = osrtp

	return request
}

// Perfil para o protocolo TLS
func ProfilerRTP_TLS_Offer(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como RTP/SAVP
	parametros.TransportProtocol = RTP_SAVP

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]ParamFlags, 0)
	osrtp := make([]OSRTP, 0)

	rtcpmux = append(rtcpmux, RTCPOffer)

	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, LoopProtect, TrustAddress)
	osrtp = append(osrtp, OSRTPAccept)

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = ICERemove
	request.DTLS = DTLSOff
	request.OSRTP = osrtp

	return request
}

// Perfil para o protocolo WS
func ProfilerRTP_WS_Offer(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como UDP/TLS/RTP/SAVP
	parametros.TransportProtocol = UDP_TLS_RTP_SAVP

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]ParamFlags, 0)
	sdes := make([]SDES, 0)

	rtcpmux = append(rtcpmux, RTCPOffer)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, LoopProtect)
	sdes = append(sdes, SDESPad)

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.SDES = sdes
	request.ICE = ICEForce
	request.DTLS = DTLSPassive

	return request
}

// Perfil para o protocolo WS
func ProfilerRTP_WSS_Offer(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como UDP/TLS/RTP/SAVPF
	parametros.TransportProtocol = UDP_TLS_RTP_SAVPF

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]ParamFlags, 0)
	sdes := make([]SDES, 0)

	rtcpmux = append(rtcpmux, RTCPOffer)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, LoopProtect, TrickleICE, TrustAddress, StrictSource, Unidirectional)
	sdes = append(sdes, SDESPad)

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.SDES = sdes
	request.ICE = ICEForce
	request.DTLS = DTLSActive

	return request
}
