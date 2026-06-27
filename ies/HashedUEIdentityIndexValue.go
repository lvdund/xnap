package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var hashedUEIdentityIndexValueConstraints = per.SizeConstraints{
	Extensible: true,
	Min:        common.Ptr(int64(13)),
	Max:        common.Ptr(int64(13)),
}

type HashedUEIdentityIndexValue struct {
	Value per.BitString
}

func (ie *HashedUEIdentityIndexValue) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, hashedUEIdentityIndexValueConstraints)
}

func (ie *HashedUEIdentityIndexValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(hashedUEIdentityIndexValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
