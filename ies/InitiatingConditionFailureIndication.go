package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	InitiatingConditionFailureIndicationChRRCReestab      = 0
	InitiatingConditionFailureIndicationChRRCSetup        = 1
	InitiatingConditionFailureIndicationChChoiceExtension = 2
)

var initiatingConditionFailureIndicationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rRCReestab"},
		{Name: "rRCSetup"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type InitiatingConditionFailureIndication struct {
	Choice          int
	RRCReestab      *RRCReestabInitiated
	RRCSetup        *RRCSetupInitiated
	ChoiceExtension []byte
}

func (ie *InitiatingConditionFailureIndication) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(initiatingConditionFailureIndicationConstraints)
	switch ie.Choice {
	case 0: // rRCReestab
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.RRCReestab.Encode(e); err != nil {
			return err
		}
	case 1: // rRCSetup
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.RRCSetup.Encode(e); err != nil {
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

func (ie *InitiatingConditionFailureIndication) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(initiatingConditionFailureIndicationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // rRCReestab
		ie.RRCReestab = new(RRCReestabInitiated)
		if err := ie.RRCReestab.Decode(d); err != nil {
			return err
		}
	case 1: // rRCSetup
		ie.RRCSetup = new(RRCSetupInitiated)
		if err := ie.RRCSetup.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
