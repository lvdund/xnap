package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var extendedUEIdentityIndexValueConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(16)),
	Max:        common.Ptr(int64(16)),
}

type ExtendedUEIdentityIndexValue struct {
	Value per.BitString
}

func (ie *ExtendedUEIdentityIndexValue) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, extendedUEIdentityIndexValueConstraints)
}

func (ie *ExtendedUEIdentityIndexValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(extendedUEIdentityIndexValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
