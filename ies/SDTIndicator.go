package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SDTIndicatorTrue int64 = 0
)

var sDTIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SDTIndicator struct {
	Value int64
}

func (ie *SDTIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sDTIndicatorConstraints)
}

func (ie *SDTIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sDTIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
