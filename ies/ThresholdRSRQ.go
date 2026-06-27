package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var thresholdRSRQConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(127)),
}

type ThresholdRSRQ struct {
	Value int64
}

func (ie *ThresholdRSRQ) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, thresholdRSRQConstraints)
}

func (ie *ThresholdRSRQ) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(thresholdRSRQConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
