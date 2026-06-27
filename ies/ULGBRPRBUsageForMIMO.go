package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uLGBRPRBUsageForMIMOConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type ULGBRPRBUsageForMIMO struct {
	Value int64
}

func (ie *ULGBRPRBUsageForMIMO) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, uLGBRPRBUsageForMIMOConstraints)
}

func (ie *ULGBRPRBUsageForMIMO) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(uLGBRPRBUsageForMIMOConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
