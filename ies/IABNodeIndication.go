package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	IABNodeIndicationTrue int64 = 0
)

var iABNodeIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type IABNodeIndication struct {
	Value int64
}

func (ie *IABNodeIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, iABNodeIndicationConstraints)
}

func (ie *IABNodeIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(iABNodeIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
