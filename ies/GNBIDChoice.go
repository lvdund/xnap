package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	GNBIDChoiceChGnbID           = 0
	GNBIDChoiceChChoiceExtension = 1
)

var gNBIDChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "gnb-ID"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type GNBIDChoice struct {
	Choice          int
	GnbID           *per.BitString
	ChoiceExtension []byte
}

func (ie *GNBIDChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(gNBIDChoiceConstraints)
	switch ie.Choice {
	case 0: // gnb-ID
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.GnbID, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(22)),
			Max:        common.Ptr(int64(32)),
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

func (ie *GNBIDChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(gNBIDChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // gnb-ID
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(22)),
			Max:        common.Ptr(int64(32)),
		})
		if err != nil {
			return err
		}
		ie.GnbID = &val
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
