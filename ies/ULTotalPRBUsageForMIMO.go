package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uLTotalPRBUsageForMIMOConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type ULTotalPRBUsageForMIMO struct {
	Value int64
}

func (ie *ULTotalPRBUsageForMIMO) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, uLTotalPRBUsageForMIMOConstraints)
}

func (ie *ULTotalPRBUsageForMIMO) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(uLTotalPRBUsageForMIMOConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
