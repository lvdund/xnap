package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TypeOfErrorNotUnderstood int64 = 0
	TypeOfErrorMissing       int64 = 1
)

var typeOfErrorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type TypeOfError struct {
	Value int64
}

func (ie *TypeOfError) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, typeOfErrorConstraints)
}

func (ie *TypeOfError) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(typeOfErrorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
