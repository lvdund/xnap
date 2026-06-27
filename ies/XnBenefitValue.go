package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var xnBenefitValueConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(8)),
}

type XnBenefitValue struct {
	Value int64
}

func (ie *XnBenefitValue) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, xnBenefitValueConstraints)
}

func (ie *XnBenefitValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(xnBenefitValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
