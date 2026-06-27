package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var expectedIdlePeriodConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(30)),
}

type ExpectedIdlePeriod struct {
	Value int64
}

func (ie *ExpectedIdlePeriod) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, expectedIdlePeriodConstraints)
}

func (ie *ExpectedIdlePeriod) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(expectedIdlePeriodConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
