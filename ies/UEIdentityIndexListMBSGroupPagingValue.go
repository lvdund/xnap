package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	UEIdentityIndexListMBSGroupPagingValueChUEIdentityIndexValueMBSGroupPaging = 0
	UEIdentityIndexListMBSGroupPagingValueChChoiceExtension                    = 1
)

var uEIdentityIndexListMBSGroupPagingValueConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "uEIdentityIndexValueMBSGroupPaging"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type UEIdentityIndexListMBSGroupPagingValue struct {
	Choice                             int
	UEIdentityIndexValueMBSGroupPaging *per.BitString
	ChoiceExtension                    []byte
}

func (ie *UEIdentityIndexListMBSGroupPagingValue) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uEIdentityIndexListMBSGroupPagingValueConstraints)
	switch ie.Choice {
	case 0: // uEIdentityIndexValueMBSGroupPaging
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.UEIdentityIndexValueMBSGroupPaging, per.SizeConstraints{
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

func (ie *UEIdentityIndexListMBSGroupPagingValue) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uEIdentityIndexListMBSGroupPagingValueConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // uEIdentityIndexValueMBSGroupPaging
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(10)),
			Max:        common.Ptr(int64(10)),
		})
		if err != nil {
			return err
		}
		ie.UEIdentityIndexValueMBSGroupPaging = &val
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
