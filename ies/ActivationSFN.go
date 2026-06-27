package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var activationSFNConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(1023)),
}

type ActivationSFN struct {
	Value int64
}

func (ie *ActivationSFN) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, activationSFNConstraints)
}

func (ie *ActivationSFN) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(activationSFNConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
