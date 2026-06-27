package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mRBIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(512)),
}

type MRBID struct {
	Value int64
}

func (ie *MRBID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, mRBIDConstraints)
}

func (ie *MRBID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(mRBIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
