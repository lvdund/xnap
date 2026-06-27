package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MDTActivationImmediateMDTOnly     int64 = 0
	MDTActivationImmediateMDTAndTrace int64 = 1
	MDTActivationLoggedMDTOnly        int64 = 2
)

var mDTActivationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type MDTActivation struct {
	Value int64
}

func (ie *MDTActivation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mDTActivationConstraints)
}

func (ie *MDTActivation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mDTActivationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
