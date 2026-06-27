package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceULNonGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type SliceULNonGBRPRBUsage struct {
	Value int64
}

func (ie *SliceULNonGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sliceULNonGBRPRBUsageConstraints)
}

func (ie *SliceULNonGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sliceULNonGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
