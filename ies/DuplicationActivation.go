package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DuplicationActivationActive   int64 = 0
	DuplicationActivationInactive int64 = 1
)

var duplicationActivationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type DuplicationActivation struct {
	Value int64
}

func (ie *DuplicationActivation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, duplicationActivationConstraints)
}

func (ie *DuplicationActivation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(duplicationActivationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
