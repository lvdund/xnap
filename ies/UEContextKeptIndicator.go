package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UEContextKeptIndicatorTrue int64 = 0
)

var uEContextKeptIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type UEContextKeptIndicator struct {
	Value int64
}

func (ie *UEContextKeptIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, uEContextKeptIndicatorConstraints)
}

func (ie *UEContextKeptIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(uEContextKeptIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
