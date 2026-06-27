package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RedundantQoSFlowIndicatorTrue  int64 = 0
	RedundantQoSFlowIndicatorFalse int64 = 1
)

var redundantQoSFlowIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RedundantQoSFlowIndicator struct {
	Value int64
}

func (ie *RedundantQoSFlowIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, redundantQoSFlowIndicatorConstraints)
}

func (ie *RedundantQoSFlowIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(redundantQoSFlowIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
