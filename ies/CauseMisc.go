package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CauseMiscControlProcessingOverload             int64 = 0
	CauseMiscHardwareFailure                       int64 = 1
	CauseMiscOAndMIntervention                     int64 = 2
	CauseMiscNotEnoughUserPlaneProcessingResources int64 = 3
	CauseMiscUnspecified                           int64 = 4
)

var causeMiscConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type CauseMisc struct {
	Value int64
}

func (ie *CauseMisc) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, causeMiscConstraints)
}

func (ie *CauseMisc) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(causeMiscConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
