package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DirectForwardingPathAvailabilityWithSourceMNDirectPathAvailable int64 = 0
)

var directForwardingPathAvailabilityWithSourceMNConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type DirectForwardingPathAvailabilityWithSourceMN struct {
	Value int64
}

func (ie *DirectForwardingPathAvailabilityWithSourceMN) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, directForwardingPathAvailabilityWithSourceMNConstraints)
}

func (ie *DirectForwardingPathAvailabilityWithSourceMN) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(directForwardingPathAvailabilityWithSourceMNConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
