package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var energyDetectionThresholdConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(-100)),
	UpperBound: common.Ptr(int64(-50)),
}

type EnergyDetectionThreshold struct {
	Value int64
}

func (ie *EnergyDetectionThreshold) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, energyDetectionThresholdConstraints)
}

func (ie *EnergyDetectionThreshold) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(energyDetectionThresholdConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
