package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var energyCostConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(10000)),
}

type EnergyCost struct {
	Value int64
}

func (ie *EnergyCost) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, energyCostConstraints)
}

func (ie *EnergyCost) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(energyCostConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
