package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	WLANMeasConfigSetup int64 = 0
)

var wLANMeasConfigConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type WLANMeasConfig struct {
	Value int64
}

func (ie *WLANMeasConfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, wLANMeasConfigConstraints)
}

func (ie *WLANMeasConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(wLANMeasConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
