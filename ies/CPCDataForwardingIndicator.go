package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CPCDataForwardingIndicatorTriggered                 int64 = 0
	CPCDataForwardingIndicatorEarlyDataTransmissionStop int64 = 1
	CPCDataForwardingIndicatorCoordinationOnly          int64 = 2
)

var cPCDataForwardingIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2},
}

type CPCDataForwardingIndicator struct {
	Value int64
}

func (ie *CPCDataForwardingIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cPCDataForwardingIndicatorConstraints)
}

func (ie *CPCDataForwardingIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cPCDataForwardingIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
