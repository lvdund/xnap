package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EUTRATransmissionBandwidthBw6   int64 = 0
	EUTRATransmissionBandwidthBw15  int64 = 1
	EUTRATransmissionBandwidthBw25  int64 = 2
	EUTRATransmissionBandwidthBw50  int64 = 3
	EUTRATransmissionBandwidthBw75  int64 = 4
	EUTRATransmissionBandwidthBw100 int64 = 5
	EUTRATransmissionBandwidthBw1   int64 = 6
)

var eUTRATransmissionBandwidthConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5},
	ExtValues:  []int64{6},
}

type EUTRATransmissionBandwidth struct {
	Value int64
}

func (ie *EUTRATransmissionBandwidth) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eUTRATransmissionBandwidthConstraints)
}

func (ie *EUTRATransmissionBandwidth) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eUTRATransmissionBandwidthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
