package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bHRLCChannelIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(16)),
	Max:        common.Ptr(int64(16)),
}

type BHRLCChannelID struct {
	Value per.BitString
}

func (ie *BHRLCChannelID) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, bHRLCChannelIDConstraints)
}

func (ie *BHRLCChannelID) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(bHRLCChannelIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
