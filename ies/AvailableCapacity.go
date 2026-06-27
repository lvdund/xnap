package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var availableCapacityConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(100)),
}

type AvailableCapacity struct {
	Value int64
}

func (ie *AvailableCapacity) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, availableCapacityConstraints)
}

func (ie *AvailableCapacity) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(availableCapacityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
