package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRPagingLongEDRXCycleInactiveHf2    int64 = 0
	NRPagingLongEDRXCycleInactiveHf4    int64 = 1
	NRPagingLongEDRXCycleInactiveHf8    int64 = 2
	NRPagingLongEDRXCycleInactiveHf16   int64 = 3
	NRPagingLongEDRXCycleInactiveHf32   int64 = 4
	NRPagingLongEDRXCycleInactiveHf64   int64 = 5
	NRPagingLongEDRXCycleInactiveHf128  int64 = 6
	NRPagingLongEDRXCycleInactiveHf256  int64 = 7
	NRPagingLongEDRXCycleInactiveHf512  int64 = 8
	NRPagingLongEDRXCycleInactiveHf1024 int64 = 9
)

var nRPagingLongEDRXCycleInactiveConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	ExtValues:  nil,
}

type NRPagingLongEDRXCycleInactive struct {
	Value int64
}

func (ie *NRPagingLongEDRXCycleInactive) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRPagingLongEDRXCycleInactiveConstraints)
}

func (ie *NRPagingLongEDRXCycleInactive) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRPagingLongEDRXCycleInactiveConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
