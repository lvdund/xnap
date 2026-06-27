package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pERScalarConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(9)),
}

type PERScalar struct {
	Value int64
}

func (ie *PERScalar) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, pERScalarConstraints)
}

func (ie *PERScalar) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(pERScalarConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
