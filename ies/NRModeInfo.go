package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRModeInfoChFdd             = 0
	NRModeInfoChTdd             = 1
	NRModeInfoChChoiceExtension = 2
)

var nRModeInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fdd"},
		{Name: "tdd"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NRModeInfo struct {
	Choice          int
	Fdd             *NRModeInfoFDD
	Tdd             *NRModeInfoTDD
	ChoiceExtension []byte
}

func (ie *NRModeInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nRModeInfoConstraints)
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

func (ie *NRModeInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nRModeInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // fdd
		ie.Fdd = new(NRModeInfoFDD)
		if err := ie.Fdd.Decode(d); err != nil {
			return err
		}
	case 1: // tdd
		ie.Tdd = new(NRModeInfoTDD)
		if err := ie.Tdd.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
