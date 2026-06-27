package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uLGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type ULGBRPRBUsage struct {
	Value int64
}

func (ie *ULGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, uLGBRPRBUsageConstraints)
}

func (ie *ULGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(uLGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
