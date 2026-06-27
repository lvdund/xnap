package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	BluetoothMeasConfigSetup int64 = 0
)

var bluetoothMeasConfigConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BluetoothMeasConfig struct {
	Value int64
}

func (ie *BluetoothMeasConfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, bluetoothMeasConfigConstraints)
}

func (ie *BluetoothMeasConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(bluetoothMeasConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
