package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DirectForwardingPathAvailabilityDirectPathAvailable int64 = 0
)

var directForwardingPathAvailabilityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type DirectForwardingPathAvailability struct {
	Value int64
}

func (ie *DirectForwardingPathAvailability) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, directForwardingPathAvailabilityConstraints)
}

func (ie *DirectForwardingPathAvailability) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(directForwardingPathAvailabilityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
