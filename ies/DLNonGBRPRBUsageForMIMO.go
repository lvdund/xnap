package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLNonGBRPRBUsageForMIMOConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type DLNonGBRPRBUsageForMIMO struct {
	Value int64
}

func (ie *DLNonGBRPRBUsageForMIMO) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dLNonGBRPRBUsageForMIMOConstraints)
}

func (ie *DLNonGBRPRBUsageForMIMO) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dLNonGBRPRBUsageForMIMOConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
