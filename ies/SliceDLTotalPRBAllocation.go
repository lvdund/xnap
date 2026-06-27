package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceDLTotalPRBAllocationConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type SliceDLTotalPRBAllocation struct {
	Value int64
}

func (ie *SliceDLTotalPRBAllocation) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sliceDLTotalPRBAllocationConstraints)
}

func (ie *SliceDLTotalPRBAllocation) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sliceDLTotalPRBAllocationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
