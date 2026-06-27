package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	SSBTransmissionBitmapChShortBitmap     = 0
	SSBTransmissionBitmapChMediumBitmap    = 1
	SSBTransmissionBitmapChLongBitmap      = 2
	SSBTransmissionBitmapChChoiceExtension = 3
)

var sSBTransmissionBitmapConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type SSBTransmissionBitmap struct {
	Choice          int
	ShortBitmap     *per.BitString
	MediumBitmap    *per.BitString
	LongBitmap      *per.BitString
	ChoiceExtension []byte
}

func (ie *SSBTransmissionBitmap) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(sSBTransmissionBitmapConstraints)
	switch ie.Choice {
	case 0: // shortBitmap
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.ShortBitmap, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(4)),
			Max:        common.Ptr(int64(4)),
		}); err != nil {
			return err
		}
	case 1: // mediumBitmap
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.MediumBitmap, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(8)),
			Max:        common.Ptr(int64(8)),
		}); err != nil {
			return err
		}
	case 2: // longBitmap
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.LongBitmap, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(64)),
			Max:        common.Ptr(int64(64)),
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

func (ie *SSBTransmissionBitmap) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(sSBTransmissionBitmapConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // shortBitmap
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(4)),
			Max:        common.Ptr(int64(4)),
		})
		if err != nil {
			return err
		}
		ie.ShortBitmap = &val
	case 1: // mediumBitmap
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(8)),
			Max:        common.Ptr(int64(8)),
		})
		if err != nil {
			return err
		}
		ie.MediumBitmap = &val
	case 2: // longBitmap
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(64)),
			Max:        common.Ptr(int64(64)),
		})
		if err != nil {
			return err
		}
		ie.LongBitmap = &val
	case 3: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
