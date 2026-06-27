package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NPNSupportChSNPN             = 0
	NPNSupportChChoiceExtensions = 1
)

var nPNSupportConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sNPN"},
		{Name: "choice-Extensions"},
	},
	ExtAlternatives: nil,
}

type NPNSupport struct {
	Choice          int
	SNPN            *NPNSupportSNPN
	ChoiceExtension []byte
}

func (ie *NPNSupport) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nPNSupportConstraints)
	switch ie.Choice {
	case 0: // sNPN
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.SNPN.Encode(e); err != nil {
			return err
		}
	case 1: // choice-Extensions
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtensions (kind=ext)
	}
	return nil
}

func (ie *NPNSupport) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nPNSupportConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // sNPN
		ie.SNPN = new(NPNSupportSNPN)
		if err := ie.SNPN.Decode(d); err != nil {
			return err
		}
	case 1: // choice-Extensions
		// TODO decode field ChoiceExtensions (kind=ext)
	}
	return nil
}
