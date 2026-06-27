package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RadioResourceStatusChNgENBRadioResourceStatus = 0
	RadioResourceStatusChGNBRadioResourceStatus   = 1
	RadioResourceStatusChChoiceExtension          = 2
)

var radioResourceStatusConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ng-eNB-RadioResourceStatus"},
		{Name: "gNB-RadioResourceStatus"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type RadioResourceStatus struct {
	Choice                   int
	NgENBRadioResourceStatus *NGENBRadioResourceStatus
	GNBRadioResourceStatus   *GNBRadioResourceStatus
	ChoiceExtension          []byte
}

func (ie *RadioResourceStatus) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(radioResourceStatusConstraints)
	switch ie.Choice {
	case 0: // ng-eNB-RadioResourceStatus
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NgENBRadioResourceStatus.Encode(e); err != nil {
			return err
		}
	case 1: // gNB-RadioResourceStatus
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.GNBRadioResourceStatus.Encode(e); err != nil {
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

func (ie *RadioResourceStatus) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(radioResourceStatusConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // ng-eNB-RadioResourceStatus
		ie.NgENBRadioResourceStatus = new(NGENBRadioResourceStatus)
		if err := ie.NgENBRadioResourceStatus.Decode(d); err != nil {
			return err
		}
	case 1: // gNB-RadioResourceStatus
		ie.GNBRadioResourceStatus = new(GNBRadioResourceStatus)
		if err := ie.GNBRadioResourceStatus.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
