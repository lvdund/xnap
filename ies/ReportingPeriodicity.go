package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReportingPeriodicityHalfThousandMs int64 = 0
	ReportingPeriodicityOneThousandMs  int64 = 1
	ReportingPeriodicityTwoThousandMs  int64 = 2
	ReportingPeriodicityFiveThousandMs int64 = 3
	ReportingPeriodicityTenThousandMs  int64 = 4
)

var reportingPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type ReportingPeriodicity struct {
	Value int64
}

func (ie *ReportingPeriodicity) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reportingPeriodicityConstraints)
}

func (ie *ReportingPeriodicity) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reportingPeriodicityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
