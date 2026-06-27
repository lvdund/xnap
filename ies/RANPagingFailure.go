package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RANPagingFailureTrue int64 = 0
)

var rANPagingFailureConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type RANPagingFailure struct {
	Value int64
}

func (ie *RANPagingFailure) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rANPagingFailureConstraints)
}

func (ie *RANPagingFailure) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rANPagingFailureConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
