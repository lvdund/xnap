package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PartialListIndicatorPartial int64 = 0
)

var partialListIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PartialListIndicator struct {
	Value int64
}

func (ie *PartialListIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, partialListIndicatorConstraints)
}

func (ie *PartialListIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(partialListIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
