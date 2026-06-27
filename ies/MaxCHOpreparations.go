package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var maxCHOpreparationsConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(8)),
}

type MaxCHOpreparations struct {
	Value int64
}

func (ie *MaxCHOpreparations) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, maxCHOpreparationsConstraints)
}

func (ie *MaxCHOpreparations) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(maxCHOpreparationsConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
