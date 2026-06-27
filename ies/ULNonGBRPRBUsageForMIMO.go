package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uLNonGBRPRBUsageForMIMOConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type ULNonGBRPRBUsageForMIMO struct {
	Value int64
}

func (ie *ULNonGBRPRBUsageForMIMO) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, uLNonGBRPRBUsageForMIMOConstraints)
}

func (ie *ULNonGBRPRBUsageForMIMO) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(uLNonGBRPRBUsageForMIMOConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
