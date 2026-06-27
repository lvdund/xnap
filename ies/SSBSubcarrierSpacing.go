package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SSBSubcarrierSpacingKHz15  int64 = 0
	SSBSubcarrierSpacingKHz30  int64 = 1
	SSBSubcarrierSpacingKHz120 int64 = 2
	SSBSubcarrierSpacingKHz240 int64 = 3
	SSBSubcarrierSpacingSpare3 int64 = 4
	SSBSubcarrierSpacingSpare2 int64 = 5
	SSBSubcarrierSpacingSpare1 int64 = 6
)

var sSBSubcarrierSpacingConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  nil,
}

type SSBSubcarrierSpacing struct {
	Value int64
}

func (ie *SSBSubcarrierSpacing) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sSBSubcarrierSpacingConstraints)
}

func (ie *SSBSubcarrierSpacing) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sSBSubcarrierSpacingConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
