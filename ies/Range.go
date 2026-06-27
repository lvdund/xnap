package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RangeM50   int64 = 0
	RangeM80   int64 = 1
	RangeM180  int64 = 2
	RangeM200  int64 = 3
	RangeM350  int64 = 4
	RangeM400  int64 = 5
	RangeM500  int64 = 6
	RangeM700  int64 = 7
	RangeM1000 int64 = 8
)

var rangeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8},
	ExtValues:  nil,
}

type Range struct {
	Value int64
}

func (ie *Range) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rangeConstraints)
}

func (ie *Range) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rangeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
