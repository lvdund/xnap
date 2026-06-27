package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ULForwardingUlForwardingProposed int64 = 0
)

var uLForwardingConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ULForwarding struct {
	Value int64
}

func (ie *ULForwarding) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, uLForwardingConstraints)
}

func (ie *ULForwarding) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(uLForwardingConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
