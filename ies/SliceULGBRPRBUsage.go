package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceULGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type SliceULGBRPRBUsage struct {
	Value int64
}

func (ie *SliceULGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sliceULGBRPRBUsageConstraints)
}

func (ie *SliceULGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sliceULGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
