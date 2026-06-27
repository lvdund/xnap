package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AvailableFastMCGRecoveryViaSRB3True int64 = 0
)

var availableFastMCGRecoveryViaSRB3Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type AvailableFastMCGRecoveryViaSRB3 struct {
	Value int64
}

func (ie *AvailableFastMCGRecoveryViaSRB3) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, availableFastMCGRecoveryViaSRB3Constraints)
}

func (ie *AvailableFastMCGRecoveryViaSRB3) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(availableFastMCGRecoveryViaSRB3Constraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
