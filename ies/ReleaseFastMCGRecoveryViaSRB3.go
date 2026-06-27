package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReleaseFastMCGRecoveryViaSRB3True int64 = 0
)

var releaseFastMCGRecoveryViaSRB3Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ReleaseFastMCGRecoveryViaSRB3 struct {
	Value int64
}

func (ie *ReleaseFastMCGRecoveryViaSRB3) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, releaseFastMCGRecoveryViaSRB3Constraints)
}

func (ie *ReleaseFastMCGRecoveryViaSRB3) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(releaseFastMCGRecoveryViaSRB3Constraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
