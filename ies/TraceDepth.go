package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TraceDepthMinimum                               int64 = 0
	TraceDepthMedium                                int64 = 1
	TraceDepthMaximum                               int64 = 2
	TraceDepthMinimumWithoutVendorSpecificExtension int64 = 3
	TraceDepthMediumWithoutVendorSpecificExtension  int64 = 4
	TraceDepthMaximumWithoutVendorSpecificExtension int64 = 5
)

var traceDepthConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5},
	ExtValues:  nil,
}

type TraceDepth struct {
	Value int64
}

func (ie *TraceDepth) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, traceDepthConstraints)
}

func (ie *TraceDepth) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(traceDepthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
