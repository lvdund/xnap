package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ECNMarkingorCongestionInformationReportingRequestChECNMarkingAtRANRequest       = 0
	ECNMarkingorCongestionInformationReportingRequestChECNMarkingAtUPFRequest       = 1
	ECNMarkingorCongestionInformationReportingRequestChCongestionInformationRequest = 2
	ECNMarkingorCongestionInformationReportingRequestChChoiceExtensions             = 3
)

var eCNMarkingorCongestionInformationReportingRequestConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eCNMarkingAtRANRequest"},
		{Name: "eCNMarkingAtUPFRequest"},
		{Name: "congestionInformationRequest"},
		{Name: "choice-Extensions"},
	},
	ExtAlternatives: nil,
}

type ECNMarkingorCongestionInformationReportingRequest struct {
	Choice                       int
	ECNMarkingAtRANRequest       *ECNMarkingAtRANRequest
	ECNMarkingAtUPFRequest       *ECNMarkingAtUPFRequest
	CongestionInformationRequest *CongestionInformationRequest
	ChoiceExtension              []byte
}

func (ie *ECNMarkingorCongestionInformationReportingRequest) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(eCNMarkingorCongestionInformationReportingRequestConstraints)
	switch ie.Choice {
	case 0: // eCNMarkingAtRANRequest
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.ECNMarkingAtRANRequest.Encode(e); err != nil {
			return err
		}
	case 1: // eCNMarkingAtUPFRequest
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ECNMarkingAtUPFRequest.Encode(e); err != nil {
			return err
		}
	case 2: // congestionInformationRequest
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.CongestionInformationRequest.Encode(e); err != nil {
			return err
		}
	case 3: // choice-Extensions
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtensions (kind=ext)
	}
	return nil
}

func (ie *ECNMarkingorCongestionInformationReportingRequest) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(eCNMarkingorCongestionInformationReportingRequestConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // eCNMarkingAtRANRequest
		ie.ECNMarkingAtRANRequest = new(ECNMarkingAtRANRequest)
		if err := ie.ECNMarkingAtRANRequest.Decode(d); err != nil {
			return err
		}
	case 1: // eCNMarkingAtUPFRequest
		ie.ECNMarkingAtUPFRequest = new(ECNMarkingAtUPFRequest)
		if err := ie.ECNMarkingAtUPFRequest.Decode(d); err != nil {
			return err
		}
	case 2: // congestionInformationRequest
		ie.CongestionInformationRequest = new(CongestionInformationRequest)
		if err := ie.CongestionInformationRequest.Decode(d); err != nil {
			return err
		}
	case 3: // choice-Extensions
		// TODO decode field ChoiceExtensions (kind=ext)
	}
	return nil
}
