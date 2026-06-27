package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceDLGBRPRBUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type SliceDLGBRPRBUsage struct {
	Value int64
}

func (ie *SliceDLGBRPRBUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sliceDLGBRPRBUsageConstraints)
}

func (ie *SliceDLGBRPRBUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sliceDLGBRPRBUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
