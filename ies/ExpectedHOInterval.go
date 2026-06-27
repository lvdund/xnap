package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ExpectedHOIntervalSec15    int64 = 0
	ExpectedHOIntervalSec30    int64 = 1
	ExpectedHOIntervalSec60    int64 = 2
	ExpectedHOIntervalSec90    int64 = 3
	ExpectedHOIntervalSec120   int64 = 4
	ExpectedHOIntervalSec180   int64 = 5
	ExpectedHOIntervalLongTime int64 = 6
)

var expectedHOIntervalConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  nil,
}

type ExpectedHOInterval struct {
	Value int64
}

func (ie *ExpectedHOInterval) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, expectedHOIntervalConstraints)
}

func (ie *ExpectedHOInterval) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(expectedHOIntervalConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
