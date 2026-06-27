package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCPACRequestInitiation int64 = 0
)

var sCPACRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SCPACRequest struct {
	Value int64
}

func (ie *SCPACRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCPACRequestConstraints)
}

func (ie *SCPACRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCPACRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
