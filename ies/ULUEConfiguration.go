package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ULUEConfigurationNoData int64 = 0
	ULUEConfigurationShared int64 = 1
	ULUEConfigurationOnly   int64 = 2
)

var uLUEConfigurationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type ULUEConfiguration struct {
	Value int64
}

func (ie *ULUEConfiguration) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, uLUEConfigurationConstraints)
}

func (ie *ULUEConfiguration) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(uLUEConfigurationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
