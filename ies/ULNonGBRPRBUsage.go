package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uLNonGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type ULNonGBRPRBUsage struct {
	Value int64
}

func (ie *ULNonGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, uLNonGBRPRBUsageConstraints)
}

func (ie *ULNonGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(uLNonGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
