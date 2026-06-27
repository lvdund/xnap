package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	HSNAFlexibleHard         int64 = 0
	HSNAFlexibleSoft         int64 = 1
	HSNAFlexibleNotavailable int64 = 2
)

var hSNAFlexibleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type HSNAFlexible struct {
	Value int64
}

func (ie *HSNAFlexible) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, hSNAFlexibleConstraints)
}

func (ie *HSNAFlexible) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(hSNAFlexibleConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
