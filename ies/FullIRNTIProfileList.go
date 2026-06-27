package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	FullIRNTIProfileListChFullIRNTIProfile0 = 0
	FullIRNTIProfileListChFullIRNTIProfile1 = 1
	FullIRNTIProfileListChFullIRNTIProfile2 = 2
	FullIRNTIProfileListChFullIRNTIProfile3 = 3
	FullIRNTIProfileListChChoiceExtension   = 4
)

var fullIRNTIProfileListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "full-I-RNTI-Profile-0"},
		{Name: "full-I-RNTI-Profile-1"},
		{Name: "full-I-RNTI-Profile-2"},
		{Name: "full-I-RNTI-Profile-3"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type FullIRNTIProfileList struct {
	Choice            int
	FullIRNTIProfile0 *per.BitString
	FullIRNTIProfile1 *per.BitString
	FullIRNTIProfile2 *per.BitString
	FullIRNTIProfile3 *per.BitString
	ChoiceExtension   []byte
}

func (ie *FullIRNTIProfileList) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(fullIRNTIProfileListConstraints)
	switch ie.Choice {
	case 0: // full-I-RNTI-Profile-0
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.FullIRNTIProfile0, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(21)),
			Max:        common.Ptr(int64(21)),
		}); err != nil {
			return err
		}
	case 1: // full-I-RNTI-Profile-1
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.FullIRNTIProfile1, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(18)),
			Max:        common.Ptr(int64(18)),
		}); err != nil {
			return err
		}
	case 2: // full-I-RNTI-Profile-2
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.FullIRNTIProfile2, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(15)),
			Max:        common.Ptr(int64(15)),
		}); err != nil {
			return err
		}
	case 3: // full-I-RNTI-Profile-3
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.FullIRNTIProfile3, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(12)),
			Max:        common.Ptr(int64(12)),
		}); err != nil {
			return err
		}
	case 4: // choice-extension
		if err := choice.EncodeChoice(4, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *FullIRNTIProfileList) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(fullIRNTIProfileListConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // full-I-RNTI-Profile-0
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(21)),
			Max:        common.Ptr(int64(21)),
		})
		if err != nil {
			return err
		}
		ie.FullIRNTIProfile0 = &val
	case 1: // full-I-RNTI-Profile-1
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(18)),
			Max:        common.Ptr(int64(18)),
		})
		if err != nil {
			return err
		}
		ie.FullIRNTIProfile1 = &val
	case 2: // full-I-RNTI-Profile-2
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(15)),
			Max:        common.Ptr(int64(15)),
		})
		if err != nil {
			return err
		}
		ie.FullIRNTIProfile2 = &val
	case 3: // full-I-RNTI-Profile-3
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(12)),
			Max:        common.Ptr(int64(12)),
		})
		if err != nil {
			return err
		}
		ie.FullIRNTIProfile3 = &val
	case 4: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
