package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SensorMeasConfigSetup int64 = 0
)

var sensorMeasConfigConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SensorMeasConfig struct {
	Value int64
}

func (ie *SensorMeasConfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sensorMeasConfigConstraints)
}

func (ie *SensorMeasConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sensorMeasConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
