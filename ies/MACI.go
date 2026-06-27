package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mACIConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(16)),
	Max:        common.Ptr(int64(16)),
}

type MACI struct {
	Value per.BitString
}

func (ie *MACI) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, mACIConstraints)
}

func (ie *MACI) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(mACIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
