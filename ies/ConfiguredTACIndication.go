package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ConfiguredTACIndicationTrue int64 = 0
)

var configuredTACIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ConfiguredTACIndication struct {
	Value int64
}

func (ie *ConfiguredTACIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, configuredTACIndicationConstraints)
}

func (ie *ConfiguredTACIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(configuredTACIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
