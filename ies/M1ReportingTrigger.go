package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M1ReportingTriggerPeriodic                 int64 = 0
	M1ReportingTriggerA2eventtriggered         int64 = 1
	M1ReportingTriggerA2eventtriggeredPeriodic int64 = 2
)

var m1ReportingTriggerConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type M1ReportingTrigger struct {
	Value int64
}

func (ie *M1ReportingTrigger) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m1ReportingTriggerConstraints)
}

func (ie *M1ReportingTrigger) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m1ReportingTriggerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
