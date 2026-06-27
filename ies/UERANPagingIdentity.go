package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	UERANPagingIdentityChIRNTIFull       = 0
	UERANPagingIdentityChChoiceExtension = 1
)

var uERANPagingIdentityConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "i-RNTI-full"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type UERANPagingIdentity struct {
	Choice          int
	IRNTIFull       *per.BitString
	ChoiceExtension []byte
}

func (ie *UERANPagingIdentity) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uERANPagingIdentityConstraints)
	switch ie.Choice {
	case 0: // i-RNTI-full
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.IRNTIFull, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(40)),
			Max:        common.Ptr(int64(40)),
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

func (ie *UERANPagingIdentity) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uERANPagingIdentityConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // i-RNTI-full
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(40)),
			Max:        common.Ptr(int64(40)),
		})
		if err != nil {
			return err
		}
		ie.IRNTIFull = &val
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
