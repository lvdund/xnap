package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bitRateConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(4000000000000)),
}

type BitRate struct {
	Value int64
}

func (ie *BitRate) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, bitRateConstraints)
}

func (ie *BitRate) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(bitRateConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
