package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReportingPeriodicityForDataCollectionHalfThousandMs int64 = 0
	ReportingPeriodicityForDataCollectionOneThousandMs  int64 = 1
	ReportingPeriodicityForDataCollectionTwoThousandMs  int64 = 2
	ReportingPeriodicityForDataCollectionFiveThousandMs int64 = 3
	ReportingPeriodicityForDataCollectionTenThousandMs  int64 = 4
)

var reportingPeriodicityForDataCollectionConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type ReportingPeriodicityForDataCollection struct {
	Value int64
}

func (ie *ReportingPeriodicityForDataCollection) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reportingPeriodicityForDataCollectionConstraints)
}

func (ie *ReportingPeriodicityForDataCollection) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reportingPeriodicityForDataCollectionConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
