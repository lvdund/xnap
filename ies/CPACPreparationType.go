package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CPACPreparationTypeSCpac int64 = 0
)

var cPACPreparationTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CPACPreparationType struct {
	Value int64
}

func (ie *CPACPreparationType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cPACPreparationTypeConstraints)
}

func (ie *CPACPreparationType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cPACPreparationTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
