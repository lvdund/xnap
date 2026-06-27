package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ServedCellInformationEUTRAModeInfoChFdd             = 0
	ServedCellInformationEUTRAModeInfoChTdd             = 1
	ServedCellInformationEUTRAModeInfoChChoiceExtension = 2
)

var servedCellInformationEUTRAModeInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fdd"},
		{Name: "tdd"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ServedCellInformationEUTRAModeInfo struct {
	Choice          int
	Fdd             *ServedCellInformationEUTRAFDDInfo
	Tdd             *ServedCellInformationEUTRATDDInfo
	ChoiceExtension []byte
}

func (ie *ServedCellInformationEUTRAModeInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(servedCellInformationEUTRAModeInfoConstraints)
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

func (ie *ServedCellInformationEUTRAModeInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(servedCellInformationEUTRAModeInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // fdd
		ie.Fdd = new(ServedCellInformationEUTRAFDDInfo)
		if err := ie.Fdd.Decode(d); err != nil {
			return err
		}
	case 1: // tdd
		ie.Tdd = new(ServedCellInformationEUTRATDDInfo)
		if err := ie.Tdd.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
