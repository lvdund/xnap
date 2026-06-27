package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLNonGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type DLNonGBRPRBUsage struct {
	Value int64
}

func (ie *DLNonGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dLNonGBRPRBUsageConstraints)
}

func (ie *DLNonGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dLNonGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
