package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pERExponentConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(9)),
}

type PERExponent struct {
	Value int64
}

func (ie *PERExponent) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, pERExponentConstraints)
}

func (ie *PERExponent) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(pERExponentConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
