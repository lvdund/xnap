package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellCapacityClassValueConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(100)),
}

type CellCapacityClassValue struct {
	Value int64
}

func (ie *CellCapacityClassValue) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cellCapacityClassValueConstraints)
}

func (ie *CellCapacityClassValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cellCapacityClassValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
