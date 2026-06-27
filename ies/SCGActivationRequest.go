package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCGActivationRequestActivateScg   int64 = 0
	SCGActivationRequestDeactivateScg int64 = 1
)

var sCGActivationRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SCGActivationRequest struct {
	Value int64
}

func (ie *SCGActivationRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCGActivationRequestConstraints)
}

func (ie *SCGActivationRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCGActivationRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
