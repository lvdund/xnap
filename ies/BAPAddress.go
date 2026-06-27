package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bAPAddressConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(10)),
	Max:        common.Ptr(int64(10)),
}

type BAPAddress struct {
	Value per.BitString
}

func (ie *BAPAddress) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, bAPAddressConstraints)
}

func (ie *BAPAddress) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(bAPAddressConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
