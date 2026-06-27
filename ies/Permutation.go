package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PermutationDfu int64 = 0
	PermutationUfd int64 = 1
)

var permutationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type Permutation struct {
	Value int64
}

func (ie *Permutation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, permutationConstraints)
}

func (ie *Permutation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(permutationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
