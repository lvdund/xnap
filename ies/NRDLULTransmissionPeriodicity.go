package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRDLULTransmissionPeriodicityMs0p5   int64 = 0
	NRDLULTransmissionPeriodicityMs0p625 int64 = 1
	NRDLULTransmissionPeriodicityMs1     int64 = 2
	NRDLULTransmissionPeriodicityMs1p25  int64 = 3
	NRDLULTransmissionPeriodicityMs2     int64 = 4
	NRDLULTransmissionPeriodicityMs2p5   int64 = 5
	NRDLULTransmissionPeriodicityMs3     int64 = 6
	NRDLULTransmissionPeriodicityMs4     int64 = 7
	NRDLULTransmissionPeriodicityMs5     int64 = 8
	NRDLULTransmissionPeriodicityMs10    int64 = 9
	NRDLULTransmissionPeriodicityMs20    int64 = 10
	NRDLULTransmissionPeriodicityMs40    int64 = 11
	NRDLULTransmissionPeriodicityMs60    int64 = 12
	NRDLULTransmissionPeriodicityMs80    int64 = 13
	NRDLULTransmissionPeriodicityMs100   int64 = 14
	NRDLULTransmissionPeriodicityMs120   int64 = 15
	NRDLULTransmissionPeriodicityMs140   int64 = 16
	NRDLULTransmissionPeriodicityMs160   int64 = 17
)

var nRDLULTransmissionPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
	ExtValues:  nil,
}

type NRDLULTransmissionPeriodicity struct {
	Value int64
}

func (ie *NRDLULTransmissionPeriodicity) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRDLULTransmissionPeriodicityConstraints)
}

func (ie *NRDLULTransmissionPeriodicity) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRDLULTransmissionPeriodicityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
