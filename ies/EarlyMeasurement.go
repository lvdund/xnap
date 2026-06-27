package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EarlyMeasurementTrue int64 = 0
)

var earlyMeasurementConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type EarlyMeasurement struct {
	Value int64
}

func (ie *EarlyMeasurement) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, earlyMeasurementConstraints)
}

func (ie *EarlyMeasurement) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(earlyMeasurementConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
