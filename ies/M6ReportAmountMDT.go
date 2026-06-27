package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M6ReportAmountMDTR1       int64 = 0
	M6ReportAmountMDTR2       int64 = 1
	M6ReportAmountMDTR4       int64 = 2
	M6ReportAmountMDTR8       int64 = 3
	M6ReportAmountMDTR16      int64 = 4
	M6ReportAmountMDTR32      int64 = 5
	M6ReportAmountMDTR64      int64 = 6
	M6ReportAmountMDTInfinity int64 = 7
)

var m6ReportAmountMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type M6ReportAmountMDT struct {
	Value int64
}

func (ie *M6ReportAmountMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m6ReportAmountMDTConstraints)
}

func (ie *M6ReportAmountMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m6ReportAmountMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
