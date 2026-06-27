package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CPCindicatorCpcInitiation   int64 = 0
	CPCindicatorCpcModification int64 = 1
	CPCindicatorCpcCancellation int64 = 2
)

var cPCindicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type CPCindicator struct {
	Value int64
}

func (ie *CPCindicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cPCindicatorConstraints)
}

func (ie *CPCindicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cPCindicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
