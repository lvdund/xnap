package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CHOtriggerChoInitiation int64 = 0
	CHOtriggerChoReplace    int64 = 1
)

var cHOtriggerConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type CHOtrigger struct {
	Value int64
}

func (ie *CHOtrigger) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cHOtriggerConstraints)
}

func (ie *CHOtrigger) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cHOtriggerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
