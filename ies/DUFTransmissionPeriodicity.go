package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DUFTransmissionPeriodicityMs0p5   int64 = 0
	DUFTransmissionPeriodicityMs0p625 int64 = 1
	DUFTransmissionPeriodicityMs1     int64 = 2
	DUFTransmissionPeriodicityMs1p25  int64 = 3
	DUFTransmissionPeriodicityMs2     int64 = 4
	DUFTransmissionPeriodicityMs2p5   int64 = 5
	DUFTransmissionPeriodicityMs5     int64 = 6
	DUFTransmissionPeriodicityMs10    int64 = 7
)

var dUFTransmissionPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type DUFTransmissionPeriodicity struct {
	Value int64
}

func (ie *DUFTransmissionPeriodicity) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dUFTransmissionPeriodicityConstraints)
}

func (ie *DUFTransmissionPeriodicity) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dUFTransmissionPeriodicityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
