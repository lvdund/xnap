package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SubcarrierSpacingKHz15  int64 = 0
	SubcarrierSpacingKHz30  int64 = 1
	SubcarrierSpacingKHz120 int64 = 2
	SubcarrierSpacingKHz240 int64 = 3
	SubcarrierSpacingSpare3 int64 = 4
	SubcarrierSpacingSpare2 int64 = 5
	SubcarrierSpacingSpare1 int64 = 6
	SubcarrierSpacingKHz60  int64 = 7
)

var subcarrierSpacingConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  []int64{7},
}

type SubcarrierSpacing struct {
	Value int64
}

func (ie *SubcarrierSpacing) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, subcarrierSpacingConstraints)
}

func (ie *SubcarrierSpacing) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(subcarrierSpacingConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
