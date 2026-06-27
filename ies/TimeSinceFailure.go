package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var timeSinceFailureConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(172800)),
}

type TimeSinceFailure struct {
	Value int64
}

func (ie *TimeSinceFailure) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, timeSinceFailureConstraints)
}

func (ie *TimeSinceFailure) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(timeSinceFailureConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
