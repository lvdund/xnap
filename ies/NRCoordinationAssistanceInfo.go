package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRCoordinationAssistanceInfoCoordinationNotRequired int64 = 0
)

var nRCoordinationAssistanceInfoConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type NRCoordinationAssistanceInfo struct {
	Value int64
}

func (ie *NRCoordinationAssistanceInfo) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRCoordinationAssistanceInfoConstraints)
}

func (ie *NRCoordinationAssistanceInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRCoordinationAssistanceInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
