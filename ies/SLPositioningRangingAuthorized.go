package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SLPositioningRangingAuthorizedAuthorized    int64 = 0
	SLPositioningRangingAuthorizedNotAuthorized int64 = 1
)

var sLPositioningRangingAuthorizedConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SLPositioningRangingAuthorized struct {
	Value int64
}

func (ie *SLPositioningRangingAuthorized) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sLPositioningRangingAuthorizedConstraints)
}

func (ie *SLPositioningRangingAuthorized) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sLPositioningRangingAuthorizedConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
