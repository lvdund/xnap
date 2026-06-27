package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCGConfigurationQueryTrue int64 = 0
)

var sCGConfigurationQueryConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SCGConfigurationQuery struct {
	Value int64
}

func (ie *SCGConfigurationQuery) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCGConfigurationQueryConstraints)
}

func (ie *SCGConfigurationQuery) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCGConfigurationQueryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
