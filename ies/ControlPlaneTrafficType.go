package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var controlPlaneTrafficTypeConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(3)),
}

type ControlPlaneTrafficType struct {
	Value int64
}

func (ie *ControlPlaneTrafficType) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, controlPlaneTrafficTypeConstraints)
}

func (ie *ControlPlaneTrafficType) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(controlPlaneTrafficTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
