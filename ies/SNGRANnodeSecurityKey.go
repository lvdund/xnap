package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNGRANnodeSecurityKeyConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(256)),
	Max:        common.Ptr(int64(256)),
}

type SNGRANnodeSecurityKey struct {
	Value per.BitString
}

func (ie *SNGRANnodeSecurityKey) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, sNGRANnodeSecurityKeyConstraints)
}

func (ie *SNGRANnodeSecurityKey) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(sNGRANnodeSecurityKeyConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
