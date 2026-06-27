package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FddOrTddChFdd             = 0
	FddOrTddChTdd             = 1
	FddOrTddChChoiceExtension = 2
)

var fddOrTddConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fdd"},
		{Name: "tdd"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type FddOrTdd struct {
	Choice          int
	Fdd             *NPRACHConfigurationFDD
	Tdd             *NPRACHConfigurationTDD
	ChoiceExtension []byte
}

func (ie *FddOrTdd) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(fddOrTddConstraints)
	switch ie.Choice {
	case 0: // fdd
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.Fdd.Encode(e); err != nil {
			return err
		}
	case 1: // tdd
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Tdd.Encode(e); err != nil {
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

func (ie *FddOrTdd) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(fddOrTddConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // fdd
		ie.Fdd = new(NPRACHConfigurationFDD)
		if err := ie.Fdd.Decode(d); err != nil {
			return err
		}
	case 1: // tdd
		ie.Tdd = new(NPRACHConfigurationTDD)
		if err := ie.Tdd.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
