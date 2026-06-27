package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AerialUEAuthorized    int64 = 0
	AerialUENotAuthorized int64 = 1
)

var aerialUEConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type AerialUE struct {
	Value int64
}

func (ie *AerialUE) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, aerialUEConstraints)
}

func (ie *AerialUE) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(aerialUEConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
