package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EUTRACoordinationAssistanceInfoCoordinationNotRequired int64 = 0
)

var eUTRACoordinationAssistanceInfoConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type EUTRACoordinationAssistanceInfo struct {
	Value int64
}

func (ie *EUTRACoordinationAssistanceInfo) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eUTRACoordinationAssistanceInfoConstraints)
}

func (ie *EUTRACoordinationAssistanceInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eUTRACoordinationAssistanceInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
