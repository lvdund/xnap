package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var fiveQIConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type FiveQI struct {
	Value int64
}

func (ie *FiveQI) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, fiveQIConstraints)
}

func (ie *FiveQI) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(fiveQIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
