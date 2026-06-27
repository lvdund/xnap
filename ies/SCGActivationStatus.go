package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCGActivationStatusScgActivated   int64 = 0
	SCGActivationStatusScgDeactivated int64 = 1
)

var sCGActivationStatusConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SCGActivationStatus struct {
	Value int64
}

func (ie *SCGActivationStatus) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCGActivationStatusConstraints)
}

func (ie *SCGActivationStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCGActivationStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
