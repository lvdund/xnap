package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CNTypeRestrictionsForServingEpcForbidden int64 = 0
)

var cNTypeRestrictionsForServingConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CNTypeRestrictionsForServing struct {
	Value int64
}

func (ie *CNTypeRestrictionsForServing) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cNTypeRestrictionsForServingConstraints)
}

func (ie *CNTypeRestrictionsForServing) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cNTypeRestrictionsForServingConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
