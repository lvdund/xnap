package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLGBRPRBUsageForMIMOConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type DLGBRPRBUsageForMIMO struct {
	Value int64
}

func (ie *DLGBRPRBUsageForMIMO) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dLGBRPRBUsageForMIMOConstraints)
}

func (ie *DLGBRPRBUsageForMIMO) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dLGBRPRBUsageForMIMOConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
