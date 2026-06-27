package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EUTRAPagingEDRXCycleHfhalf int64 = 0
	EUTRAPagingEDRXCycleHf1    int64 = 1
	EUTRAPagingEDRXCycleHf2    int64 = 2
	EUTRAPagingEDRXCycleHf4    int64 = 3
	EUTRAPagingEDRXCycleHf6    int64 = 4
	EUTRAPagingEDRXCycleHf8    int64 = 5
	EUTRAPagingEDRXCycleHf10   int64 = 6
	EUTRAPagingEDRXCycleHf12   int64 = 7
	EUTRAPagingEDRXCycleHf14   int64 = 8
	EUTRAPagingEDRXCycleHf16   int64 = 9
	EUTRAPagingEDRXCycleHf32   int64 = 10
	EUTRAPagingEDRXCycleHf64   int64 = 11
	EUTRAPagingEDRXCycleHf128  int64 = 12
	EUTRAPagingEDRXCycleHf256  int64 = 13
)

var eUTRAPagingEDRXCycleConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	ExtValues:  nil,
}

type EUTRAPagingEDRXCycle struct {
	Value int64
}

func (ie *EUTRAPagingEDRXCycle) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eUTRAPagingEDRXCycleConstraints)
}

func (ie *EUTRAPagingEDRXCycle) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eUTRAPagingEDRXCycleConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
