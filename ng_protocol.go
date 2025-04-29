package rtpengine

import "fmt"

type ParametrosOption func(c *RequestRtp) error

// Gera oferta do SDP com passagem de Parametros
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

// Gera Atendimendo do SDP com passagem de Parametros
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

// Gera Delete da sessão no rtpengine com passagem de Parametros
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

// Adcionar um lista de flags para rtpengine
func (c *RequestRtp) SetFlags(flags []ParamFlags) ParametrosOption {
	return func(s *RequestRtp) error {
		s.ParamsOptStringArray.Flags = append(s.ParamsOptStringArray.Flags, flags...)
		return nil
	}
}

// Manipular o Transport Protocol do SDP
func (c *RequestRtp) SetTransportProtocol(proto TransportProtocol) ParametrosOption {
	return func(s *RequestRtp) error {
		s.TransportProtocol = proto
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

// Manipular codecs marca quais serão aceito na lista do SDP
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

// Manipular codecs remover da lista do SDP
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

// Bloquear todos os codecs, exceto aqueles fornecidos na lista de permissões.
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

// Desabilitar a criptografia SDES na oferta
func (c *RequestRtp) DesabilitarSDES() ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]SDES, 0)
		sdes = append(sdes, SDESOff)
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// Excluir pacotes de criptografia individuais
func (c *RequestRtp) DeletesSDES(cript []CryptoSuite) ParametrosOption {
	return func(s *RequestRtp) error {
		sdes := make([]SDES, 0)
		for _, o := range cript {
			sdes = append(sdes, "no-"+SDES(o))
		}
		s.ParamsOptStringArray.SDES = append(s.ParamsOptStringArray.SDES, sdes...)
		return nil
	}
}

// Permitir apenas o pacotes de criptografia individuais
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

// Manipulador de atributos do SDP suporta adicionar, remover e substituir
func (c *RequestRtp) SetAttrChange(sdpAttr *ParamsSdpAttrSections) ParametrosOption {
	return func(s *RequestRtp) error {
		s.SdpAttr = sdpAttr
		return nil
	}
}

// Manipulador de atributos do SDP suporta adicionar, remover e substituir
func (c *RequestRtp) SetViaBranchTag(branch string) ParametrosOption {
	return func(s *RequestRtp) error {
		s.ViaBranch = branch
		return nil
	}
}

// Adicionar o valor de ptime do codec no offer valor a ser utilizado e inteiro
func (c *RequestRtp) SetPtimeCodecOffer(ptime int) ParametrosOption {
	return func(s *RequestRtp) error {
		s.Ptime = ptime
		return nil
	}
}

// Adicionar o valor de ptime do codec no answer valor a ser utilizado e inteiro
func (c *RequestRtp) SetPtimeCodecAnswer(ptime int) ParametrosOption {
	return func(s *RequestRtp) error {
		s.PtimeReverse = ptime
		return nil
	}
}

// Adicionar o received-from Usado se os endereços SDP não forem confiáveis
func (c *RequestRtp) SetReceivedFrom(addressFamily AddressFamily, Address string) ParametrosOption {
	return func(s *RequestRtp) error {
		receivedFrom := make([]string, 0)
		s.ReceivedFrom = append(receivedFrom, string(addressFamily), Address)
		return nil
	}
}

func (c *RequestRtp) SetMediaAddress(Address string) ParametrosOption {
	return func(s *RequestRtp) error {
		s.MediaAddress = Address
		return nil
	}
}

func (c *RequestRtp) RecordYes() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "yes"
		return nil
	}
}

func (c *RequestRtp) RecordNo() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "no"
		return nil
	}
}

func (c *RequestRtp) RecordOn() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "on"
		return nil
	}
}

func (c *RequestRtp) RecordOff() ParametrosOption {
	return func(s *RequestRtp) error {
		s.RecordCall = "off"
		return nil
	}
}
