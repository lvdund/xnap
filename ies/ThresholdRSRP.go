package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var thresholdRSRPConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(127)),
}

type ThresholdRSRP struct {
	Value int64
}

func (ie *ThresholdRSRP) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, thresholdRSRPConstraints)
}

func (ie *ThresholdRSRP) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(thresholdRSRPConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
