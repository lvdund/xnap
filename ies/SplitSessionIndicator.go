package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SplitSessionIndicatorSplit int64 = 0
)

var splitSessionIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SplitSessionIndicator struct {
	Value int64
}

func (ie *SplitSessionIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, splitSessionIndicatorConstraints)
}

func (ie *SplitSessionIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(splitSessionIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
