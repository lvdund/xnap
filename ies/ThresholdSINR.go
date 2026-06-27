package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var thresholdSINRConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(127)),
}

type ThresholdSINR struct {
	Value int64
}

func (ie *ThresholdSINR) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, thresholdSINRConstraints)
}

func (ie *ThresholdSINR) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(thresholdSINRConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
