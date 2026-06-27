package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReflectiveQoSAttributeSubjectToReflectiveQoS int64 = 0
)

var reflectiveQoSAttributeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ReflectiveQoSAttribute struct {
	Value int64
}

func (ie *ReflectiveQoSAttribute) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reflectiveQoSAttributeConstraints)
}

func (ie *ReflectiveQoSAttribute) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reflectiveQoSAttributeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
