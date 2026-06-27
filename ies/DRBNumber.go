package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBNumberConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(32)),
}

type DRBNumber struct {
	Value int64
}

func (ie *DRBNumber) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dRBNumberConstraints)
}

func (ie *DRBNumber) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dRBNumberConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
