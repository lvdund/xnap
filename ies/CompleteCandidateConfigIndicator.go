package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CompleteCandidateConfigIndicatorCompleteCandidateConfig int64 = 0
)

var completeCandidateConfigIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CompleteCandidateConfigIndicator struct {
	Value int64
}

func (ie *CompleteCandidateConfigIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, completeCandidateConfigIndicatorConstraints)
}

func (ie *CompleteCandidateConfigIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(completeCandidateConfigIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
