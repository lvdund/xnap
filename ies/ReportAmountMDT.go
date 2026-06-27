package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReportAmountMDTR1       int64 = 0
	ReportAmountMDTR2       int64 = 1
	ReportAmountMDTR4       int64 = 2
	ReportAmountMDTR8       int64 = 3
	ReportAmountMDTR16      int64 = 4
	ReportAmountMDTR32      int64 = 5
	ReportAmountMDTR64      int64 = 6
	ReportAmountMDTInfinity int64 = 7
)

var reportAmountMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type ReportAmountMDT struct {
	Value int64
}

func (ie *ReportAmountMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reportAmountMDTConstraints)
}

func (ie *ReportAmountMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reportAmountMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
