package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UserPlaneErrorIndicatorGtpuErrorIndicationReceived int64 = 0
)

var userPlaneErrorIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type UserPlaneErrorIndicator struct {
	Value int64
}

func (ie *UserPlaneErrorIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, userPlaneErrorIndicatorConstraints)
}

func (ie *UserPlaneErrorIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(userPlaneErrorIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
