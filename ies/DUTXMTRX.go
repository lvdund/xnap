package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DUTXMTRXSupported            int64 = 0
	DUTXMTRXNotSupported         int64 = 1
	DUTXMTRXSupportedFDMRequired int64 = 2
)

var dUTXMTRXConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type DUTXMTRX struct {
	Value int64
}

func (ie *DUTXMTRX) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dUTXMTRXConstraints)
}

func (ie *DUTXMTRX) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dUTXMTRXConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
