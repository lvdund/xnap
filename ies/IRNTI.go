package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	IRNTIChIRNTIFull       = 0
	IRNTIChIRNTIShort      = 1
	IRNTIChChoiceExtension = 2
)

var iRNTIConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "i-RNTI-full"},
		{Name: "i-RNTI-short"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type IRNTI struct {
	Choice          int
	IRNTIFull       *per.BitString
	IRNTIShort      *per.BitString
	ChoiceExtension []byte
}

func (ie *IRNTI) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(iRNTIConstraints)
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
	case 1: // i-RNTI-short
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.IRNTIShort, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(24)),
			Max:        common.Ptr(int64(24)),
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

func (ie *IRNTI) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(iRNTIConstraints)
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
	case 1: // i-RNTI-short
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(24)),
			Max:        common.Ptr(int64(24)),
		})
		if err != nil {
			return err
		}
		ie.IRNTIShort = &val
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
