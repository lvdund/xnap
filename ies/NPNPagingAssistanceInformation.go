package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NPNPagingAssistanceInformationChPniNpnInformation = 0
	NPNPagingAssistanceInformationChChoiceExtension   = 1
)

var nPNPagingAssistanceInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pni-npn-Information"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NPNPagingAssistanceInformation struct {
	Choice            int
	PniNpnInformation *NPNPagingAssistanceInformationPNINPN
	ChoiceExtension   []byte
}

func (ie *NPNPagingAssistanceInformation) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nPNPagingAssistanceInformationConstraints)
	switch ie.Choice {
	case 0: // pni-npn-Information
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.PniNpnInformation.Encode(e); err != nil {
			return err
		}
	case 1: // choice-extension
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *NPNPagingAssistanceInformation) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nPNPagingAssistanceInformationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // pni-npn-Information
		ie.PniNpnInformation = new(NPNPagingAssistanceInformationPNINPN)
		if err := ie.PniNpnInformation.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
