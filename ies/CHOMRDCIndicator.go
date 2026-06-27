package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CHOMRDCIndicatorTrue             int64 = 0
	CHOMRDCIndicatorCoordinationOnly int64 = 1
)

var cHOMRDCIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  []int64{1},
}

type CHOMRDCIndicator struct {
	Value int64
}

func (ie *CHOMRDCIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cHOMRDCIndicatorConstraints)
}

func (ie *CHOMRDCIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cHOMRDCIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
