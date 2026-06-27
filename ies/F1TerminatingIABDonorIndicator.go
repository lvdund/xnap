package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	F1TerminatingIABDonorIndicatorTrue int64 = 0
)

var f1TerminatingIABDonorIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type F1TerminatingIABDonorIndicator struct {
	Value int64
}

func (ie *F1TerminatingIABDonorIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, f1TerminatingIABDonorIndicatorConstraints)
}

func (ie *F1TerminatingIABDonorIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(f1TerminatingIABDonorIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
