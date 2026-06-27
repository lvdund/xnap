package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qosMonitoringReportingFrequencyConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(1800)),
}

type QosMonitoringReportingFrequency struct {
	Value int64
}

func (ie *QosMonitoringReportingFrequency) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, qosMonitoringReportingFrequencyConstraints)
}

func (ie *QosMonitoringReportingFrequency) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(qosMonitoringReportingFrequencyConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
