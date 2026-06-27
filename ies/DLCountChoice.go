package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DLCountChoiceChCount12bits     = 0
	DLCountChoiceChCount18bits     = 1
	DLCountChoiceChChoiceExtension = 2
)

var dLCountChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "count12bits"},
		{Name: "count18bits"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type DLCountChoice struct {
	Choice          int
	Count12bits     *COUNTPDCPSN12
	Count18bits     *COUNTPDCPSN18
	ChoiceExtension []byte
}

func (ie *DLCountChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(dLCountChoiceConstraints)
	switch ie.Choice {
	case 0: // count12bits
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.Count12bits.Encode(e); err != nil {
			return err
		}
	case 1: // count18bits
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Count18bits.Encode(e); err != nil {
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

func (ie *DLCountChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(dLCountChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // count12bits
		ie.Count12bits = new(COUNTPDCPSN12)
		if err := ie.Count12bits.Decode(d); err != nil {
			return err
		}
	case 1: // count18bits
		ie.Count18bits = new(COUNTPDCPSN18)
		if err := ie.Count18bits.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
