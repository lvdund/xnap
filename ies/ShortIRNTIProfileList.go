package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ShortIRNTIProfileListChShortIRNTIProfile0 = 0
	ShortIRNTIProfileListChShortIRNTIProfile1 = 1
	ShortIRNTIProfileListChChoiceExtension    = 2
)

var shortIRNTIProfileListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "short-I-RNTI-Profile-0"},
		{Name: "short-I-RNTI-Profile-1"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ShortIRNTIProfileList struct {
	Choice             int
	ShortIRNTIProfile0 *per.BitString
	ShortIRNTIProfile1 *per.BitString
	ChoiceExtension    []byte
}

func (ie *ShortIRNTIProfileList) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(shortIRNTIProfileListConstraints)
	switch ie.Choice {
	case 0: // short-I-RNTI-Profile-0
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.ShortIRNTIProfile0, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(8)),
			Max:        common.Ptr(int64(8)),
		}); err != nil {
			return err
		}
	case 1: // short-I-RNTI-Profile-1
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.ShortIRNTIProfile1, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		}); err != nil {
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

func (ie *ShortIRNTIProfileList) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(shortIRNTIProfileListConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // short-I-RNTI-Profile-0
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(8)),
			Max:        common.Ptr(int64(8)),
		})
		if err != nil {
			return err
		}
		ie.ShortIRNTIProfile0 = &val
	case 1: // short-I-RNTI-Profile-1
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		})
		if err != nil {
			return err
		}
		ie.ShortIRNTIProfile1 = &val
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
