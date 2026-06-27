package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MTSDTIndicatorTrue int64 = 0
)

var mTSDTIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MTSDTIndicator struct {
	Value int64
}

func (ie *MTSDTIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mTSDTIndicatorConstraints)
}

func (ie *MTSDTIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mTSDTIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
