package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	MBSFNSubframeAllocationEUTRAChOneframe        = 0
	MBSFNSubframeAllocationEUTRAChFourframes      = 1
	MBSFNSubframeAllocationEUTRAChChoiceExtension = 2
)

var mBSFNSubframeAllocationEUTRAConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneframe"},
		{Name: "fourframes"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type MBSFNSubframeAllocationEUTRA struct {
	Choice          int
	Oneframe        *per.BitString
	Fourframes      *per.BitString
	ChoiceExtension []byte
}

func (ie *MBSFNSubframeAllocationEUTRA) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(mBSFNSubframeAllocationEUTRAConstraints)
	switch ie.Choice {
	case 0: // oneframe
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.Oneframe, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		}); err != nil {
			return err
		}
	case 1: // fourframes
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.Fourframes, per.SizeConstraints{
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

func (ie *MBSFNSubframeAllocationEUTRA) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(mBSFNSubframeAllocationEUTRAConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // oneframe
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		})
		if err != nil {
			return err
		}
		ie.Oneframe = &val
	case 1: // fourframes
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(24)),
			Max:        common.Ptr(int64(24)),
		})
		if err != nil {
			return err
		}
		ie.Fourframes = &val
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
