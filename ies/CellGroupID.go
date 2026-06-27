package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellGroupIDConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(common.MaxnoofSCellGroups)),
}

type CellGroupID struct {
	Value int64
}

func (ie *CellGroupID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cellGroupIDConstraints)
}

func (ie *CellGroupID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cellGroupIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
