package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var expectedActivityPeriodConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(30)),
}

type ExpectedActivityPeriod struct {
	Value int64
}

func (ie *ExpectedActivityPeriod) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, expectedActivityPeriodConstraints)
}

func (ie *ExpectedActivityPeriod) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(expectedActivityPeriodConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
