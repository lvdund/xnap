package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RequestedFastMCGRecoveryViaSRB3ReleaseTrue int64 = 0
)

var requestedFastMCGRecoveryViaSRB3ReleaseConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type RequestedFastMCGRecoveryViaSRB3Release struct {
	Value int64
}

func (ie *RequestedFastMCGRecoveryViaSRB3Release) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, requestedFastMCGRecoveryViaSRB3ReleaseConstraints)
}

func (ie *RequestedFastMCGRecoveryViaSRB3Release) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(requestedFastMCGRecoveryViaSRB3ReleaseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
