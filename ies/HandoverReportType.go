package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	HandoverReportTypeHoTooEarly          int64 = 0
	HandoverReportTypeHoToWrongCell       int64 = 1
	HandoverReportTypeIntersystempingpong int64 = 2
)

var handoverReportTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type HandoverReportType struct {
	Value int64
}

func (ie *HandoverReportType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, handoverReportTypeConstraints)
}

func (ie *HandoverReportType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(handoverReportTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
