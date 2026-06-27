package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CyclicPrefixEUTRAULNormal   int64 = 0
	CyclicPrefixEUTRAULExtended int64 = 1
)

var cyclicPrefixEUTRAULConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type CyclicPrefixEUTRAUL struct {
	Value int64
}

func (ie *CyclicPrefixEUTRAUL) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cyclicPrefixEUTRAULConstraints)
}

func (ie *CyclicPrefixEUTRAUL) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cyclicPrefixEUTRAULConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
