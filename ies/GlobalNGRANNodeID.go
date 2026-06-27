package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	GlobalNGRANNodeIDChGNB             = 0
	GlobalNGRANNodeIDChNgENB           = 1
	GlobalNGRANNodeIDChChoiceExtension = 2
)

var globalNGRANNodeIDConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "gNB"},
		{Name: "ng-eNB"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type GlobalNGRANNodeID struct {
	Choice          int
	GNB             *GlobalgNBID
	NgENB           *GlobalngeNBID
	ChoiceExtension []byte
}

func (ie *GlobalNGRANNodeID) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(globalNGRANNodeIDConstraints)
	switch ie.Choice {
	case 0: // gNB
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.GNB.Encode(e); err != nil {
			return err
		}
	case 1: // ng-eNB
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.NgENB.Encode(e); err != nil {
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

func (ie *GlobalNGRANNodeID) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(globalNGRANNodeIDConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // gNB
		ie.GNB = new(GlobalgNBID)
		if err := ie.GNB.Decode(d); err != nil {
			return err
		}
	case 1: // ng-eNB
		ie.NgENB = new(GlobalngeNBID)
		if err := ie.NgENB.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
