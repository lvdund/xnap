package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var lCIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(32)),
}

type LCID struct {
	Value int64
}

func (ie *LCID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, lCIDConstraints)
}

func (ie *LCID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(lCIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
