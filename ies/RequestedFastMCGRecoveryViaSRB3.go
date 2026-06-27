package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RequestedFastMCGRecoveryViaSRB3True int64 = 0
)

var requestedFastMCGRecoveryViaSRB3Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type RequestedFastMCGRecoveryViaSRB3 struct {
	Value int64
}

func (ie *RequestedFastMCGRecoveryViaSRB3) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, requestedFastMCGRecoveryViaSRB3Constraints)
}

func (ie *RequestedFastMCGRecoveryViaSRB3) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(requestedFastMCGRecoveryViaSRB3Constraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
