package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MDTAlignmentInfoChSBasedMDT       = 0
	MDTAlignmentInfoChChoiceExtension = 1
)

var mDTAlignmentInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "s-BasedMDT"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type MDTAlignmentInfo struct {
	Choice          int
	SBasedMDT       *SBasedMDT
	ChoiceExtension []byte
}

func (ie *MDTAlignmentInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(mDTAlignmentInfoConstraints)
	switch ie.Choice {
	case 0: // s-BasedMDT
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.SBasedMDT.Encode(e); err != nil {
			return err
		}
	case 1: // choice-extension
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *MDTAlignmentInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(mDTAlignmentInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // s-BasedMDT
		ie.SBasedMDT = new(SBasedMDT)
		if err := ie.SBasedMDT.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
