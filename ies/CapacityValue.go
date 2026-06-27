package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var capacityValueConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type CapacityValue struct {
	Value int64
}

func (ie *CapacityValue) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, capacityValueConstraints)
}

func (ie *CapacityValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(capacityValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
