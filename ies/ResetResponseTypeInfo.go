package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ResetResponseTypeInfoChFullReset       = 0
	ResetResponseTypeInfoChPartialReset    = 1
	ResetResponseTypeInfoChChoiceExtension = 2
)

var resetResponseTypeInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fullReset"},
		{Name: "partialReset"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ResetResponseTypeInfo struct {
	Choice          int
	FullReset       *ResetResponseTypeInfoFull
	PartialReset    *ResetResponseTypeInfoPartial
	ChoiceExtension []byte
}

func (ie *ResetResponseTypeInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(resetResponseTypeInfoConstraints)
	switch ie.Choice {
	case 0: // fullReset
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.FullReset.Encode(e); err != nil {
			return err
		}
	case 1: // partialReset
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.PartialReset.Encode(e); err != nil {
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

func (ie *ResetResponseTypeInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(resetResponseTypeInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // fullReset
		ie.FullReset = new(ResetResponseTypeInfoFull)
		if err := ie.FullReset.Decode(d); err != nil {
			return err
		}
	case 1: // partialReset
		ie.PartialReset = new(ResetResponseTypeInfoPartial)
		if err := ie.PartialReset.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
