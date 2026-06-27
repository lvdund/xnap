package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var hysteresisConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(30)),
}

type Hysteresis struct {
	Value int64
}

func (ie *Hysteresis) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, hysteresisConstraints)
}

func (ie *Hysteresis) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(hysteresisConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
