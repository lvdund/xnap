package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M5ReportAmountMDTR1       int64 = 0
	M5ReportAmountMDTR2       int64 = 1
	M5ReportAmountMDTR4       int64 = 2
	M5ReportAmountMDTR8       int64 = 3
	M5ReportAmountMDTR16      int64 = 4
	M5ReportAmountMDTR32      int64 = 5
	M5ReportAmountMDTR64      int64 = 6
	M5ReportAmountMDTInfinity int64 = 7
)

var m5ReportAmountMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type M5ReportAmountMDT struct {
	Value int64
}

func (ie *M5ReportAmountMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m5ReportAmountMDTConstraints)
}

func (ie *M5ReportAmountMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m5ReportAmountMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
