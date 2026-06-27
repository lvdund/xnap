package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(44)),
	Max:        common.Ptr(int64(44)),
}

type NID struct {
	Value per.BitString
}

func (ie *NID) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, nIDConstraints)
}

func (ie *NID) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(nIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
