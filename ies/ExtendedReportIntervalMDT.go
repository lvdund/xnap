package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ExtendedReportIntervalMDTMs20480 int64 = 0
	ExtendedReportIntervalMDTMs40960 int64 = 1
)

var extendedReportIntervalMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type ExtendedReportIntervalMDT struct {
	Value int64
}

func (ie *ExtendedReportIntervalMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, extendedReportIntervalMDTConstraints)
}

func (ie *ExtendedReportIntervalMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(extendedReportIntervalMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
