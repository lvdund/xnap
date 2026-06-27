package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type DLGBRPRBUsage struct {
	Value int64
}

func (ie *DLGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dLGBRPRBUsageConstraints)
}

func (ie *DLGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dLGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
