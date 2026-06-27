package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NonUPTrafficChNonUPTrafficType        = 0
	NonUPTrafficChControlPlaneTrafficType = 1
	NonUPTrafficChChoiceExtension         = 2
)

var nonUPTrafficConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nonUPTrafficType"},
		{Name: "controlPlaneTrafficType"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NonUPTraffic struct {
	Choice                  int
	NonUPTrafficType        *NonUPTrafficType
	ControlPlaneTrafficType *ControlPlaneTrafficType
	ChoiceExtension         []byte
}

func (ie *NonUPTraffic) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nonUPTrafficConstraints)
	switch ie.Choice {
	case 0: // nonUPTrafficType
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NonUPTrafficType.Encode(e); err != nil {
			return err
		}
	case 1: // controlPlaneTrafficType
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ControlPlaneTrafficType.Encode(e); err != nil {
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

func (ie *NonUPTraffic) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nonUPTrafficConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nonUPTrafficType
		ie.NonUPTrafficType = new(NonUPTrafficType)
		if err := ie.NonUPTrafficType.Decode(d); err != nil {
			return err
		}
	case 1: // controlPlaneTrafficType
		ie.ControlPlaneTrafficType = new(ControlPlaneTrafficType)
		if err := ie.ControlPlaneTrafficType.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
