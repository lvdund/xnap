package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var activationIDforCellActivationConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type ActivationIDforCellActivation struct {
	Value int64
}

func (ie *ActivationIDforCellActivation) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, activationIDforCellActivationConstraints)
}

func (ie *ActivationIDforCellActivation) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(activationIDforCellActivationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
