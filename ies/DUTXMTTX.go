package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DUTXMTTXSupported            int64 = 0
	DUTXMTTXNotSupported         int64 = 1
	DUTXMTTXSupportedFDMRequired int64 = 2
)

var dUTXMTTXConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type DUTXMTTX struct {
	Value int64
}

func (ie *DUTXMTTX) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dUTXMTTXConstraints)
}

func (ie *DUTXMTTX) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dUTXMTTXConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
