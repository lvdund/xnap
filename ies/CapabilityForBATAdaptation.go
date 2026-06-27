package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CapabilityForBATAdaptationTrue int64 = 0
)

var capabilityForBATAdaptationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CapabilityForBATAdaptation struct {
	Value int64
}

func (ie *CapabilityForBATAdaptation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, capabilityForBATAdaptationConstraints)
}

func (ie *CapabilityForBATAdaptation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(capabilityForBATAdaptationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
