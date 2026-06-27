package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var offeredCapacityConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(16777216)),
}

type OfferedCapacity struct {
	Value int64
}

func (ie *OfferedCapacity) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, offeredCapacityConstraints)
}

func (ie *OfferedCapacity) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(offeredCapacityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
