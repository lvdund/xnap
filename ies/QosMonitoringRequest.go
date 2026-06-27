package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QosMonitoringRequestUl   int64 = 0
	QosMonitoringRequestDl   int64 = 1
	QosMonitoringRequestBoth int64 = 2
)

var qosMonitoringRequestConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type QosMonitoringRequest struct {
	Value int64
}

func (ie *QosMonitoringRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qosMonitoringRequestConstraints)
}

func (ie *QosMonitoringRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qosMonitoringRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
