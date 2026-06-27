package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRPagingEDRXCycleInactiveHfquarter int64 = 0
	NRPagingEDRXCycleInactiveHfhalf    int64 = 1
	NRPagingEDRXCycleInactiveHf1       int64 = 2
)

var nRPagingEDRXCycleInactiveConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type NRPagingEDRXCycleInactive struct {
	Value int64
}

func (ie *NRPagingEDRXCycleInactive) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRPagingEDRXCycleInactiveConstraints)
}

func (ie *NRPagingEDRXCycleInactive) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRPagingEDRXCycleInactiveConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
