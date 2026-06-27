package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ENBIDChoiceChEnbIDMacro      = 0
	ENBIDChoiceChEnbIDShortmacro = 1
	ENBIDChoiceChEnbIDLongmacro  = 2
	ENBIDChoiceChChoiceExtension = 3
)

var eNBIDChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "enb-ID-macro"},
		{Name: "enb-ID-shortmacro"},
		{Name: "enb-ID-longmacro"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ENBIDChoice struct {
	Choice          int
	EnbIDMacro      *per.BitString
	EnbIDShortmacro *per.BitString
	EnbIDLongmacro  *per.BitString
	ChoiceExtension []byte
}

func (ie *ENBIDChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(eNBIDChoiceConstraints)
	switch ie.Choice {
	case 0: // enb-ID-macro
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.EnbIDMacro, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(20)),
			Max:        common.Ptr(int64(20)),
		}); err != nil {
			return err
		}
	case 1: // enb-ID-shortmacro
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.EnbIDShortmacro, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(18)),
			Max:        common.Ptr(int64(18)),
		}); err != nil {
			return err
		}
	case 2: // enb-ID-longmacro
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.EnbIDLongmacro, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(21)),
			Max:        common.Ptr(int64(21)),
		}); err != nil {
			return err
		}
	case 3: // choice-extension
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *ENBIDChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(eNBIDChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // enb-ID-macro
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(20)),
			Max:        common.Ptr(int64(20)),
		})
		if err != nil {
			return err
		}
		ie.EnbIDMacro = &val
	case 1: // enb-ID-shortmacro
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(18)),
			Max:        common.Ptr(int64(18)),
		})
		if err != nil {
			return err
		}
		ie.EnbIDShortmacro = &val
	case 2: // enb-ID-longmacro
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(21)),
			Max:        common.Ptr(int64(21)),
		})
		if err != nil {
			return err
		}
		ie.EnbIDLongmacro = &val
	case 3: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
