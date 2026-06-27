package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TargetCGIChNr              = 0
	TargetCGIChEUtra           = 1
	TargetCGIChChoiceExtension = 2
)

var targetCGIConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr"},
		{Name: "e-utra"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type TargetCGI struct {
	Choice          int
	Nr              *NRCGI
	EUtra           *EUTRACGI
	ChoiceExtension []byte
}

func (ie *TargetCGI) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(targetCGIConstraints)
	switch ie.Choice {
	case 0: // nr
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.Nr.Encode(e); err != nil {
			return err
		}
	case 1: // e-utra
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.EUtra.Encode(e); err != nil {
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

func (ie *TargetCGI) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(targetCGIConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nr
		ie.Nr = new(NRCGI)
		if err := ie.Nr.Decode(d); err != nil {
			return err
		}
	case 1: // e-utra
		ie.EUtra = new(EUTRACGI)
		if err := ie.EUtra.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
