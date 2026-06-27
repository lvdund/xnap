package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mDTLocationInfoConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(8)),
	Max:        common.Ptr(int64(8)),
}

type MDTLocationInfo struct {
	Value per.BitString
}

func (ie *MDTLocationInfo) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, mDTLocationInfoConstraints)
}

func (ie *MDTLocationInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(mDTLocationInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
