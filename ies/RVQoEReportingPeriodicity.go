package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RVQoEReportingPeriodicityMs120  int64 = 0
	RVQoEReportingPeriodicityMs240  int64 = 1
	RVQoEReportingPeriodicityMs480  int64 = 2
	RVQoEReportingPeriodicityMs640  int64 = 3
	RVQoEReportingPeriodicityMs1024 int64 = 4
)

var rVQoEReportingPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type RVQoEReportingPeriodicity struct {
	Value int64
}

func (ie *RVQoEReportingPeriodicity) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rVQoEReportingPeriodicityConstraints)
}

func (ie *RVQoEReportingPeriodicity) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rVQoEReportingPeriodicityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
