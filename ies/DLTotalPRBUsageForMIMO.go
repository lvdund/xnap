package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLTotalPRBUsageForMIMOConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type DLTotalPRBUsageForMIMO struct {
	Value int64
}

func (ie *DLTotalPRBUsageForMIMO) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dLTotalPRBUsageForMIMOConstraints)
}

func (ie *DLTotalPRBUsageForMIMO) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dLTotalPRBUsageForMIMOConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
