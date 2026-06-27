package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uLTotalPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type ULTotalPRBUsage struct {
	Value int64
}

func (ie *ULTotalPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, uLTotalPRBUsageConstraints)
}

func (ie *ULTotalPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(uLTotalPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
