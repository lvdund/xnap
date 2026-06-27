package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AerialControllerUEAuthorized    int64 = 0
	AerialControllerUENotAuthorized int64 = 1
)

var aerialControllerUEConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type AerialControllerUE struct {
	Value int64
}

func (ie *AerialControllerUE) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, aerialControllerUEConstraints)
}

func (ie *AerialControllerUE) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(aerialControllerUEConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
