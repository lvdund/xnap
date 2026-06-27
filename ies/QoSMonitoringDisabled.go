package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoSMonitoringDisabledTrue int64 = 0
)

var qoSMonitoringDisabledConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type QoSMonitoringDisabled struct {
	Value int64
}

func (ie *QoSMonitoringDisabled) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoSMonitoringDisabledConstraints)
}

func (ie *QoSMonitoringDisabled) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoSMonitoringDisabledConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
