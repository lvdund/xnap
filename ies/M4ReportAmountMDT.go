package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M4ReportAmountMDTR1       int64 = 0
	M4ReportAmountMDTR2       int64 = 1
	M4ReportAmountMDTR4       int64 = 2
	M4ReportAmountMDTR8       int64 = 3
	M4ReportAmountMDTR16      int64 = 4
	M4ReportAmountMDTR32      int64 = 5
	M4ReportAmountMDTR64      int64 = 6
	M4ReportAmountMDTInfinity int64 = 7
)

var m4ReportAmountMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type M4ReportAmountMDT struct {
	Value int64
}

func (ie *M4ReportAmountMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m4ReportAmountMDTConstraints)
}

func (ie *M4ReportAmountMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m4ReportAmountMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
