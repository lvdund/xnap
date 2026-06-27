package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NPNMobilityInformationChSnpnMobilityInformation   = 0
	NPNMobilityInformationChPniNpnMobilityInformation = 1
	NPNMobilityInformationChChoiceExtension           = 2
)

var nPNMobilityInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "snpn-mobility-information"},
		{Name: "pni-npn-mobility-information"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NPNMobilityInformation struct {
	Choice                    int
	SnpnMobilityInformation   *NPNMobilityInformationSNPN
	PniNpnMobilityInformation *NPNMobilityInformationPNINPN
	ChoiceExtension           []byte
}

func (ie *NPNMobilityInformation) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nPNMobilityInformationConstraints)
	switch ie.Choice {
	case 0: // snpn-mobility-information
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.SnpnMobilityInformation.Encode(e); err != nil {
			return err
		}
	case 1: // pni-npn-mobility-information
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.PniNpnMobilityInformation.Encode(e); err != nil {
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

func (ie *NPNMobilityInformation) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nPNMobilityInformationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // snpn-mobility-information
		ie.SnpnMobilityInformation = new(NPNMobilityInformationSNPN)
		if err := ie.SnpnMobilityInformation.Decode(d); err != nil {
			return err
		}
	case 1: // pni-npn-mobility-information
		ie.PniNpnMobilityInformation = new(NPNMobilityInformationPNINPN)
		if err := ie.PniNpnMobilityInformation.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
