package rtpengine

import (
	"fmt"
)

func SDPOffering(parametros *ParamsOptString, options ...ParametrosOption) (*RequestRtp, error) {
	request := &RequestRtp{
		Command:               fmt.Sprint(Offer),
		ParamsOptString:       parametros,
		ParamsOptInt:          &ParamsOptInt{},
		ParamsOptStringArray:  &ParamsOptStringArray{},
		ParamsSdpAttrSections: &ParamsSdpAttrSections{},
		ParamsSdpAttrCommands: &ParamsSdpAttrCommands{},
	}

	for _, o := range options {
		if err := o(request); err != nil {
			return nil, err
		}
	}
	return request, nil
}

func ProfilerRTP_UDP(command string, parametros *ParamsOptString) *RequestRtp {
	request := &RequestRtp{
		Command:               command,
		ParamsOptString:       parametros,
		ParamsOptInt:          &ParamsOptInt{},
		ParamsOptStringArray:  &ParamsOptStringArray{},
		ParamsSdpAttrSections: &ParamsSdpAttrSections{},
		ParamsSdpAttrCommands: &ParamsSdpAttrCommands{},
	}

	// definir o protocolo com RTP/AVP
	parametros.TransportProtocol = string(RTP_AVP)

	rtcpmux := make([]ParamRTCPMux, 0)
	replace := make([]ParamReplace, 0)
	flags := make([]string, 0)

	rtcpmux = append(rtcpmux, RTCP_Demux)
	replace = append(replace, SessionConnection, Origin)
	flags = append(flags, string(StripExtmap), string(NoRtcpAttribute))

	request.RtcpMux = rtcpmux
	request.Replace = replace
	request.Flags = flags
	request.ICE = string(ICERemove)
	request.DTLS = string(DTLSOff)

	return request
}
