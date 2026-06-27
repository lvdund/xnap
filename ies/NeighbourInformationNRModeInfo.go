package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NeighbourInformationNRModeInfoChFddInfo         = 0
	NeighbourInformationNRModeInfoChTddInfo         = 1
	NeighbourInformationNRModeInfoChChoiceExtension = 2
)

var neighbourInformationNRModeInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fdd-info"},
		{Name: "tdd-info"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NeighbourInformationNRModeInfo struct {
	Choice          int
	FddInfo         *NeighbourInformationNRModeFDDInfo
	TddInfo         *NeighbourInformationNRModeTDDInfo
	ChoiceExtension []byte
}

func (ie *NeighbourInformationNRModeInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(neighbourInformationNRModeInfoConstraints)
	switch ie.Choice {
	case 0: // fdd-info
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.FddInfo.Encode(e); err != nil {
			return err
		}
	case 1: // tdd-info
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.TddInfo.Encode(e); err != nil {
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

func (ie *NeighbourInformationNRModeInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(neighbourInformationNRModeInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // fdd-info
		ie.FddInfo = new(NeighbourInformationNRModeFDDInfo)
		if err := ie.FddInfo.Decode(d); err != nil {
			return err
		}
	case 1: // tdd-info
		ie.TddInfo = new(NeighbourInformationNRModeTDDInfo)
		if err := ie.TddInfo.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
