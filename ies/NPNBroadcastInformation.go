package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NPNBroadcastInformationChSnpnInformation   = 0
	NPNBroadcastInformationChPniNpnInformation = 1
	NPNBroadcastInformationChChoiceExtension   = 2
)

var nPNBroadcastInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "snpn-Information"},
		{Name: "pni-npn-Information"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NPNBroadcastInformation struct {
	Choice            int
	SnpnInformation   *NPNBroadcastInformationSNPN
	PniNpnInformation *NPNBroadcastInformationPNINPN
	ChoiceExtension   []byte
}

func (ie *NPNBroadcastInformation) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nPNBroadcastInformationConstraints)
	switch ie.Choice {
	case 0: // snpn-Information
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.SnpnInformation.Encode(e); err != nil {
			return err
		}
	case 1: // pni-npn-Information
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.PniNpnInformation.Encode(e); err != nil {
			return err
		}
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *NPNBroadcastInformation) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nPNBroadcastInformationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // snpn-Information
		ie.SnpnInformation = new(NPNBroadcastInformationSNPN)
		if err := ie.SnpnInformation.Decode(d); err != nil {
			return err
		}
	case 1: // pni-npn-Information
		ie.PniNpnInformation = new(NPNBroadcastInformationPNINPN)
		if err := ie.PniNpnInformation.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
