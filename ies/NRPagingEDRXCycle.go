package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRPagingEDRXCycleHfquarter int64 = 0
	NRPagingEDRXCycleHfhalf    int64 = 1
	NRPagingEDRXCycleHf1       int64 = 2
	NRPagingEDRXCycleHf2       int64 = 3
	NRPagingEDRXCycleHf4       int64 = 4
	NRPagingEDRXCycleHf8       int64 = 5
	NRPagingEDRXCycleHf16      int64 = 6
	NRPagingEDRXCycleHf32      int64 = 7
	NRPagingEDRXCycleHf64      int64 = 8
	NRPagingEDRXCycleHf128     int64 = 9
	NRPagingEDRXCycleHf256     int64 = 10
	NRPagingEDRXCycleHf512     int64 = 11
	NRPagingEDRXCycleHf1024    int64 = 12
)

var nRPagingEDRXCycleConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	ExtValues:  nil,
}

type NRPagingEDRXCycle struct {
	Value int64
}

func (ie *NRPagingEDRXCycle) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRPagingEDRXCycleConstraints)
}

func (ie *NRPagingEDRXCycle) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRPagingEDRXCycleConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
