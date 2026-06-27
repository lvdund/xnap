package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ExpectedUEMobilityStationary int64 = 0
	ExpectedUEMobilityMobile     int64 = 1
)

var expectedUEMobilityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type ExpectedUEMobility struct {
	Value int64
}

func (ie *ExpectedUEMobility) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, expectedUEMobilityConstraints)
}

func (ie *ExpectedUEMobility) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(expectedUEMobilityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
