package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PedestrianUEAuthorized    int64 = 0
	PedestrianUENotAuthorized int64 = 1
)

var pedestrianUEConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PedestrianUE struct {
	Value int64
}

func (ie *PedestrianUE) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pedestrianUEConstraints)
}

func (ie *PedestrianUE) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pedestrianUEConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
