package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SplitSRBsTypesSrb1     int64 = 0
	SplitSRBsTypesSrb2     int64 = 1
	SplitSRBsTypesSrb1and2 int64 = 2
)

var splitSRBsTypesConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type SplitSRBsTypes struct {
	Value int64
}

func (ie *SplitSRBsTypes) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, splitSRBsTypesConstraints)
}

func (ie *SplitSRBsTypes) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(splitSRBsTypesConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
