package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DURXMTRXSupported            int64 = 0
	DURXMTRXNotSupported         int64 = 1
	DURXMTRXSupportedFDMRequired int64 = 2
)

var dURXMTRXConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type DURXMTRX struct {
	Value int64
}

func (ie *DURXMTRX) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dURXMTRXConstraints)
}

func (ie *DURXMTRX) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dURXMTRXConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
