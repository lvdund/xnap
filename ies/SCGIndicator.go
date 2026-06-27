package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCGIndicatorReleased int64 = 0
)

var sCGIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SCGIndicator struct {
	Value int64
}

func (ie *SCGIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCGIndicatorConstraints)
}

func (ie *SCGIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCGIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
