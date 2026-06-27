package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	VehicleUEAuthorized    int64 = 0
	VehicleUENotAuthorized int64 = 1
)

var vehicleUEConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type VehicleUE struct {
	Value int64
}

func (ie *VehicleUE) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, vehicleUEConstraints)
}

func (ie *VehicleUE) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(vehicleUEConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
