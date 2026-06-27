package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	BeamMeasurementIndicationM1True int64 = 0
)

var beamMeasurementIndicationM1Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BeamMeasurementIndicationM1 struct {
	Value int64
}

func (ie *BeamMeasurementIndicationM1) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, beamMeasurementIndicationM1Constraints)
}

func (ie *BeamMeasurementIndicationM1) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(beamMeasurementIndicationM1Constraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
