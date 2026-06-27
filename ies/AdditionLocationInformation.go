package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AdditionLocationInformationIncludePSCell int64 = 0
)

var additionLocationInformationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type AdditionLocationInformation struct {
	Value int64
}

func (ie *AdditionLocationInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, additionLocationInformationConstraints)
}

func (ie *AdditionLocationInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(additionLocationInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
