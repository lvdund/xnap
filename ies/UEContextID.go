package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UEContextIDChRRCResume           = 0
	UEContextIDChRRRCReestablishment = 1
	UEContextIDChChoiceExtension     = 2
)

var uEContextIDConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rRCResume"},
		{Name: "rRRCReestablishment"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type UEContextID struct {
	Choice              int
	RRCResume           *UEContextIDforRRCResume
	RRRCReestablishment *UEContextIDforRRCReestablishment
	ChoiceExtension     []byte
}

func (ie *UEContextID) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uEContextIDConstraints)
	switch ie.Choice {
	case 0: // rRCResume
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.RRCResume.Encode(e); err != nil {
			return err
		}
	case 1: // rRRCReestablishment
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.RRRCReestablishment.Encode(e); err != nil {
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

func (ie *UEContextID) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uEContextIDConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // rRCResume
		ie.RRCResume = new(UEContextIDforRRCResume)
		if err := ie.RRCResume.Decode(d); err != nil {
			return err
		}
	case 1: // rRRCReestablishment
		ie.RRRCReestablishment = new(UEContextIDforRRCReestablishment)
		if err := ie.RRRCReestablishment.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
