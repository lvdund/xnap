package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SPRAvailabilitySprAvailable int64 = 0
)

var sPRAvailabilityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SPRAvailability struct {
	Value int64
}

func (ie *SPRAvailability) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sPRAvailabilityConstraints)
}

func (ie *SPRAvailability) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sPRAvailabilityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
