package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UESpecificDRXV32  int64 = 0
	UESpecificDRXV64  int64 = 1
	UESpecificDRXV128 int64 = 2
	UESpecificDRXV256 int64 = 3
)

var uESpecificDRXConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  nil,
}

type UESpecificDRX struct {
	Value int64
}

func (ie *UESpecificDRX) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, uESpecificDRXConstraints)
}

func (ie *UESpecificDRX) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(uESpecificDRXConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
