package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MaxIPrateBitrate64kbs int64 = 0
	MaxIPrateMaxUErate    int64 = 1
)

var maxIPrateConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MaxIPrate struct {
	Value int64
}

func (ie *MaxIPrate) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, maxIPrateConstraints)
}

func (ie *MaxIPrate) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(maxIPrateConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
