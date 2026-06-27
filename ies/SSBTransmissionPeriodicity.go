package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SSBTransmissionPeriodicitySf10  int64 = 0
	SSBTransmissionPeriodicitySf20  int64 = 1
	SSBTransmissionPeriodicitySf40  int64 = 2
	SSBTransmissionPeriodicitySf80  int64 = 3
	SSBTransmissionPeriodicitySf160 int64 = 4
	SSBTransmissionPeriodicitySf320 int64 = 5
	SSBTransmissionPeriodicitySf640 int64 = 6
	SSBTransmissionPeriodicitySf5   int64 = 7
)

var sSBTransmissionPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  []int64{7},
}

type SSBTransmissionPeriodicity struct {
	Value int64
}

func (ie *SSBTransmissionPeriodicity) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sSBTransmissionPeriodicityConstraints)
}

func (ie *SSBTransmissionPeriodicity) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sSBTransmissionPeriodicityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
