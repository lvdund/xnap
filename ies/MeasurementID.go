package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var measurementIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(4095)),
}

type MeasurementID struct {
	Value int64
}

func (ie *MeasurementID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, measurementIDConstraints)
}

func (ie *MeasurementID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(measurementIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
