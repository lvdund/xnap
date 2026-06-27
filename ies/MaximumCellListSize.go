package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var maximumCellListSizeConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(16384)),
}

type MaximumCellListSize struct {
	Value int64
}

func (ie *MaximumCellListSize) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, maximumCellListSizeConstraints)
}

func (ie *MaximumCellListSize) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(maximumCellListSizeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
