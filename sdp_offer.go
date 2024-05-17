package rtpengine

import (
	"fmt"
)

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

// Perfil para o protocolo UDP
func ProfilerRTP_UDP(command string, parametros *ParamsOptString) *RequestRtp {
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
	flags := make([]string, 0)
	sdes := make([]string, 0)

	rtcpmux = append(rtcpmux, RTCPDemux)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, string(StripExtmap), string(NoRtcpAttribute))
	sdes = append(sdes, string(SDESOff))

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = string(ICERemove)
	request.DTLS = string(DTLSOff)
	request.SDES = sdes

	return request
}

// Perfil para o protocolo TCP
func ProfilerRTP_TCP(command string, parametros *ParamsOptString) *RequestRtp {
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
	flags := make([]string, 0)
	osrtp := make([]string, 0)

	rtcpmux = append(rtcpmux, RTCPDemux)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, string(LoopProtect), string(StrictSource))
	osrtp = append(osrtp, string(OSRTPOffer))

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = string(ICERemove)
	request.DTLS = string(DTLSOff)
	request.OSRTP = osrtp

	return request
}

// Perfil para o protocolo TLS
func ProfilerRTP_TLS(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:              command,
		ParamsOptString:      parametros,
		ParamsOptInt:         &ParamsOptInt{},
		ParamsOptStringArray: &ParamsOptStringArray{},
	}

	// definir o protocolo como RTP/SAVP
	parametros.TransportProtocol = string(RTP_SAVP)

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]string, 0)
	osrtp := make([]string, 0)

	rtcpmux = append(rtcpmux, RTCPOffer)

	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, string(LoopProtect), string(TrustAddress))
	osrtp = append(osrtp, string(OSRTPAccept))

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = string(ICERemove)
	request.DTLS = string(DTLSOff)
	request.OSRTP = osrtp

	request.SdpAttr.Global.Add = append(request.SdpAttr.Global.Add, "Software Switches CBX4")

	return request
}
