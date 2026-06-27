package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	XRBcastInformationTrue int64 = 0
)

var xRBcastInformationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type XRBcastInformation struct {
	Value int64
}

func (ie *XRBcastInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, xRBcastInformationConstraints)
}

func (ie *XRBcastInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(xRBcastInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
