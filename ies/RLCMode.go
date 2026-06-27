package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RLCModeRlcAm                 int64 = 0
	RLCModeRlcUmBidirectional    int64 = 1
	RLCModeRlcUmUnidirectionalUl int64 = 2
	RLCModeRlcUmUnidirectionalDl int64 = 3
)

var rLCModeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  nil,
}

type RLCMode struct {
	Value int64
}

func (ie *RLCMode) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rLCModeConstraints)
}

func (ie *RLCMode) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rLCModeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
