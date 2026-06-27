package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	HSNATransmissionPeriodicityMs0p5   int64 = 0
	HSNATransmissionPeriodicityMs0p625 int64 = 1
	HSNATransmissionPeriodicityMs1     int64 = 2
	HSNATransmissionPeriodicityMs1p25  int64 = 3
	HSNATransmissionPeriodicityMs2     int64 = 4
	HSNATransmissionPeriodicityMs2p5   int64 = 5
	HSNATransmissionPeriodicityMs5     int64 = 6
	HSNATransmissionPeriodicityMs10    int64 = 7
	HSNATransmissionPeriodicityMs20    int64 = 8
	HSNATransmissionPeriodicityMs40    int64 = 9
	HSNATransmissionPeriodicityMs80    int64 = 10
	HSNATransmissionPeriodicityMs160   int64 = 11
)

var hSNATransmissionPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	ExtValues:  nil,
}

type HSNATransmissionPeriodicity struct {
	Value int64
}

func (ie *HSNATransmissionPeriodicity) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, hSNATransmissionPeriodicityConstraints)
}

func (ie *HSNATransmissionPeriodicity) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(hSNATransmissionPeriodicityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
