package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRCyclicPrefixNormal   int64 = 0
	NRCyclicPrefixExtended int64 = 1
)

var nRCyclicPrefixConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type NRCyclicPrefix struct {
	Value int64
}

func (ie *NRCyclicPrefix) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRCyclicPrefixConstraints)
}

func (ie *NRCyclicPrefix) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRCyclicPrefixConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
