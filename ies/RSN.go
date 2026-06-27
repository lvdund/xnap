package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RSNV1 int64 = 0
	RSNV2 int64 = 1
)

var rSNConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RSN struct {
	Value int64
}

func (ie *RSN) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rSNConstraints)
}

func (ie *RSN) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rSNConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
