package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceDLNonGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type SliceDLNonGBRPRBUsage struct {
	Value int64
}

func (ie *SliceDLNonGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sliceDLNonGBRPRBUsageConstraints)
}

func (ie *SliceDLNonGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sliceDLNonGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
