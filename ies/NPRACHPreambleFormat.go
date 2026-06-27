package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NPRACHPreambleFormatFmt0  int64 = 0
	NPRACHPreambleFormatFmt1  int64 = 1
	NPRACHPreambleFormatFmt2  int64 = 2
	NPRACHPreambleFormatFmt0a int64 = 3
	NPRACHPreambleFormatFmt1a int64 = 4
)

var nPRACHPreambleFormatConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type NPRACHPreambleFormat struct {
	Value int64
}

func (ie *NPRACHPreambleFormat) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nPRACHPreambleFormatConstraints)
}

func (ie *NPRACHPreambleFormat) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nPRACHPreambleFormatConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
