package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AllTrafficIndicationTrue int64 = 0
)

var allTrafficIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type AllTrafficIndication struct {
	Value int64
}

func (ie *AllTrafficIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, allTrafficIndicationConstraints)
}

func (ie *AllTrafficIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(allTrafficIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
