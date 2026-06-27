package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCPACReferenceConfigRequestRequest int64 = 0
)

var sCPACReferenceConfigRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SCPACReferenceConfigRequest struct {
	Value int64
}

func (ie *SCPACReferenceConfigRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCPACReferenceConfigRequestConstraints)
}

func (ie *SCPACReferenceConfigRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCPACReferenceConfigRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
