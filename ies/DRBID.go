package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(32)),
}

type DRBID struct {
	Value int64
}

func (ie *DRBID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dRBIDConstraints)
}

func (ie *DRBID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dRBIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
