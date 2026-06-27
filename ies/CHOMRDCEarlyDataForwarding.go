package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CHOMRDCEarlyDataForwardingStop int64 = 0
)

var cHOMRDCEarlyDataForwardingConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CHOMRDCEarlyDataForwarding struct {
	Value int64
}

func (ie *CHOMRDCEarlyDataForwarding) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cHOMRDCEarlyDataForwardingConstraints)
}

func (ie *CHOMRDCEarlyDataForwarding) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cHOMRDCEarlyDataForwardingConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
