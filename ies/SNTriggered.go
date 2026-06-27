package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SNTriggeredTrue int64 = 0
)

var sNTriggeredConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SNTriggered struct {
	Value int64
}

func (ie *SNTriggered) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNTriggeredConstraints)
}

func (ie *SNTriggered) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNTriggeredConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
