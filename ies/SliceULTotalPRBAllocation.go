package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceULTotalPRBAllocationConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type SliceULTotalPRBAllocation struct {
	Value int64
}

func (ie *SliceULTotalPRBAllocation) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sliceULTotalPRBAllocationConstraints)
}

func (ie *SliceULTotalPRBAllocation) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sliceULTotalPRBAllocationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
