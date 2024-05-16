package rtpengine

type ParametrosOption func(c *RequestRtp) error

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
			trascoder = append(trascoder, string("codec-transcode-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, trascoder...)
		return nil
	}
}

// Manipular codecs marca quais serão aceito na lista do SDP
func (c *RequestRtp) SetCodecMask(codecs []string) ParametrosOption {
	return func(s *RequestRtp) error {
		mask := make([]string, 0)
		for _, o := range codecs {
			mask = append(mask, string("codec-mask-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, mask...)
		return nil
	}
}

// Manipular codecs remover da lista do SDP
func (c *RequestRtp) SetCodecStrip(codecs []string) ParametrosOption {
	return func(s *RequestRtp) error {
		strip := make([]string, 0)
		for _, o := range codecs {
			strip = append(strip, string("codec-strip-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, strip...)
		return nil
	}
}

// Bloquear todos os codecs, exceto aqueles fornecidos na lista de permissões.
func (c *RequestRtp) SetCodecExcept(codecs []string) ParametrosOption {
	return func(s *RequestRtp) error {
		except := make([]string, 0)
		for _, o := range codecs {
			except = append(except, string("codec-except-"+o))
		}

		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, except...)
		return nil
	}
}

// Desabilitar a criptografia SDES na oferta
func (c *RequestRtp) DesabilitarSDES() ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]string, 0)
		sdes = append(sdes, string("SDES-off"))
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// Excluir pacotes de criptografia individuais
func (c *RequestRtp) DeletesSDES(cript []string) ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]string, 0)
		for _, o := range cript {
			sdes = append(sdes, string("no-"+o))
		}
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// Permitir apenas o pacotes de criptografia individuais
func (c *RequestRtp) EnableSDES(cript []string) ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]string, 0)
		for _, o := range cript {
			sdes = append(sdes, string("only-"+o))
		}
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// Qualquer atributos do ICE será removido do corpo do SDP
func (c *RequestRtp) ICERemove() ParametrosOption {
	return func(s *RequestRtp) error {
		s.ICE = "remove"
		return nil
	}
}

// Os atributos do ICE são primeiro removidos e, em seguida, novos atributos são gerados e inseridos, o que deixa o proxy de mídia como o único candidato ao ICE
func (c *RequestRtp) ICEForce() ParametrosOption {
	return func(s *RequestRtp) error {
		s.ICE = "force"
		return nil
	}
}
