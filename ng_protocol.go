package rtpengine

import "fmt"

type ParametrosOption func(c *RequestRtp) error

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

func SDPAnswer(parametros *ParamsOptString, options ...*ParamsOptStringArray) (string error) {
	return nil
}

func SDPDelete(parametros *ParamsOptString) (string error) {
	return nil
}

// Adcionar um lista de flags para rtpengine
func (c *RequestRtp) SetFlags(flags []string) ParametrosOption {
	return func(s *RequestRtp) error {
		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, flags...)
		return nil
	}
}

// Manipular o Transport Protocol do SDP
func (c *RequestRtp) SetTransportProtocol(proto TransportProtocol) ParametrosOption {
	return func(s *RequestRtp) error {
		s.TransportProtocol = string(proto)
		return nil
	}
}

// Adiciona flags de manipulação
func (c *RequestRtp) SetReplace(replace []ParamReplace) ParametrosOption {
	return func(s *RequestRtp) error {
		s.Replace = replace
		return nil
	}
}

// Manipular o comportamento do rtcp-mux
func (c *RequestRtp) SetRtcpMux(rtcpmux []ParamRTCPMux) ParametrosOption {
	return func(s *RequestRtp) error {
		s.RtcpMux = rtcpmux
		return nil
	}
}

// Manipular o transcoder dos codecs
func (c *RequestRtp) SetCodecEncoder(codecs []string) ParametrosOption {
	return func(s *RequestRtp) error {
		trascoder := make([]string, 0)
		for _, o := range codecs {
			trascoder = append(trascoder, string("codec-set-transcode="+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, trascoder...)
		return nil
	}
}
