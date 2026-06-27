package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M7ReportAmountMDTR1       int64 = 0
	M7ReportAmountMDTR2       int64 = 1
	M7ReportAmountMDTR4       int64 = 2
	M7ReportAmountMDTR8       int64 = 3
	M7ReportAmountMDTR16      int64 = 4
	M7ReportAmountMDTR32      int64 = 5
	M7ReportAmountMDTR64      int64 = 6
	M7ReportAmountMDTInfinity int64 = 7
)

var m7ReportAmountMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type M7ReportAmountMDT struct {
	Value int64
}

func (ie *M7ReportAmountMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m7ReportAmountMDTConstraints)
}

func (ie *M7ReportAmountMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m7ReportAmountMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
