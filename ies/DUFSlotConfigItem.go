package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DUFSlotConfigItemChExplicitFormat  = 0
	DUFSlotConfigItemChImplicitFormat  = 1
	DUFSlotConfigItemChChoiceExtension = 2
)

var dUFSlotConfigItemConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "explicitFormat"},
		{Name: "implicitFormat"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type DUFSlotConfigItem struct {
	Choice          int
	ExplicitFormat  *ExplicitFormat
	ImplicitFormat  *ImplicitFormat
	ChoiceExtension []byte
}

func (ie *DUFSlotConfigItem) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(dUFSlotConfigItemConstraints)
	switch ie.Choice {
	case 0: // explicitFormat
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.ExplicitFormat.Encode(e); err != nil {
			return err
		}
	case 1: // implicitFormat
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ImplicitFormat.Encode(e); err != nil {
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

func (ie *DUFSlotConfigItem) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(dUFSlotConfigItemConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // explicitFormat
		ie.ExplicitFormat = new(ExplicitFormat)
		if err := ie.ExplicitFormat.Decode(d); err != nil {
			return err
		}
	case 1: // implicitFormat
		ie.ImplicitFormat = new(ImplicitFormat)
		if err := ie.ImplicitFormat.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
