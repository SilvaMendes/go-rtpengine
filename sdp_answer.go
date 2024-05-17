package rtpengine

func ProfilerRTP_UDP_Answer(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como RTP/AVP
	parametros.TransportProtocol = string(RTP_AVP)

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]ParamFlags, 0)
	sdes := make([]SDES, 0)

	rtcpmux = append(rtcpmux, RTCPDemux)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, StripExtmap, NoRtcpAttribute)
	sdes = append(sdes, SDESPrefer)

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = ICERemove
	request.DTLS = DTLSOff
	request.SDES = sdes

	return request
}
