package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	UEIdentityIndexValueChIndexLength10   = 0
	UEIdentityIndexValueChChoiceExtension = 1
)

var uEIdentityIndexValueConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "indexLength10"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type UEIdentityIndexValue struct {
	Choice          int
	IndexLength10   *per.BitString
	ChoiceExtension []byte
}

func (ie *UEIdentityIndexValue) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uEIdentityIndexValueConstraints)
	switch ie.Choice {
	case 0: // indexLength10
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.IndexLength10, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(10)),
			Max:        common.Ptr(int64(10)),
		}); err != nil {
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

func (ie *UEIdentityIndexValue) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uEIdentityIndexValueConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // indexLength10
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(10)),
			Max:        common.Ptr(int64(10)),
		})
		if err != nil {
			return err
		}
		ie.IndexLength10 = &val
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
