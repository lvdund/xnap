package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReestablishmentIndicationReestablished int64 = 0
)

var reestablishmentIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ReestablishmentIndication struct {
	Value int64
}

func (ie *ReestablishmentIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reestablishmentIndicationConstraints)
}

func (ie *ReestablishmentIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reestablishmentIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
