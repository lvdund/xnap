package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MDTModeNRChImmediateMDT = 0
	MDTModeNRChLoggedMDT    = 1
	MDTModeNRExtension      = -1
)

var mDTModeNRConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "immediateMDT"},
		{Name: "loggedMDT"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "mDTMode-NR-Extension"},
	},
}

type MDTModeNR struct {
	Choice          int
	ImmediateMDT    *ImmediateMDTNR
	LoggedMDT       *LoggedMDTNR
	ChoiceExtension []byte
}

func (ie *MDTModeNR) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(mDTModeNRConstraints)
	switch ie.Choice {
	case 0: // immediateMDT
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.ImmediateMDT.Encode(e); err != nil {
			return err
		}
	case 1: // loggedMDT
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.LoggedMDT.Encode(e); err != nil {
			return err
		}
	default: // extension
		if err := choice.EncodeChoice(0, true, ie.ChoiceExtension); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MDTModeNR) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(mDTModeNRConstraints)
	idx, isExt, extBytes, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	if isExt {
		ie.Choice = MDTModeNRExtension
		ie.ChoiceExtension = extBytes
		return nil
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // immediateMDT
		ie.ImmediateMDT = new(ImmediateMDTNR)
		if err := ie.ImmediateMDT.Decode(d); err != nil {
			return err
		}
	case 1: // loggedMDT
		ie.LoggedMDT = new(LoggedMDTNR)
		if err := ie.LoggedMDT.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
