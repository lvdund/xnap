package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dUFSlotformatIndexConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(254)),
}

type DUFSlotformatIndex struct {
	Value int64
}

func (ie *DUFSlotformatIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dUFSlotformatIndexConstraints)
}

func (ie *DUFSlotformatIndex) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dUFSlotformatIndexConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
