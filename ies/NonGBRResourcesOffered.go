package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NonGBRResourcesOfferedTrue int64 = 0
)

var nonGBRResourcesOfferedConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type NonGBRResourcesOffered struct {
	Value int64
}

func (ie *NonGBRResourcesOffered) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nonGBRResourcesOfferedConstraints)
}

func (ie *NonGBRResourcesOffered) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nonGBRResourcesOfferedConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
