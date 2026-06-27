package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sRBIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(4)),
}

type SRBID struct {
	Value int64
}

func (ie *SRBID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sRBIDConstraints)
}

func (ie *SRBID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sRBIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
