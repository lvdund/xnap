package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DefaultDRBAllowedTrue  int64 = 0
	DefaultDRBAllowedFalse int64 = 1
)

var defaultDRBAllowedConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type DefaultDRBAllowed struct {
	Value int64
}

func (ie *DefaultDRBAllowed) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, defaultDRBAllowedConstraints)
}

func (ie *DefaultDRBAllowed) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(defaultDRBAllowedConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
