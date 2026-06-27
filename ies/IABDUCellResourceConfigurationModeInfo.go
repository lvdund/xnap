package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	IABDUCellResourceConfigurationModeInfoChTDD             = 0
	IABDUCellResourceConfigurationModeInfoChFDD             = 1
	IABDUCellResourceConfigurationModeInfoChChoiceExtension = 2
)

var iABDUCellResourceConfigurationModeInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "tDD"},
		{Name: "fDD"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type IABDUCellResourceConfigurationModeInfo struct {
	Choice          int
	TDD             *IABDUCellResourceConfigurationTDDInfo
	FDD             *IABDUCellResourceConfigurationFDDInfo
	ChoiceExtension []byte
}

func (ie *IABDUCellResourceConfigurationModeInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(iABDUCellResourceConfigurationModeInfoConstraints)
	switch ie.Choice {
	case 0: // tDD
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.TDD.Encode(e); err != nil {
			return err
		}
	case 1: // fDD
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.FDD.Encode(e); err != nil {
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

func (ie *IABDUCellResourceConfigurationModeInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(iABDUCellResourceConfigurationModeInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // tDD
		ie.TDD = new(IABDUCellResourceConfigurationTDDInfo)
		if err := ie.TDD.Decode(d); err != nil {
			return err
		}
	case 1: // fDD
		ie.FDD = new(IABDUCellResourceConfigurationFDDInfo)
		if err := ie.FDD.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
