package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	BarringExemptionforEmerCallInfoTrue int64 = 0
)

var barringExemptionforEmerCallInfoConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BarringExemptionforEmerCallInfo struct {
	Value int64
}

func (ie *BarringExemptionforEmerCallInfo) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, barringExemptionforEmerCallInfoConstraints)
}

func (ie *BarringExemptionforEmerCallInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(barringExemptionforEmerCallInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
