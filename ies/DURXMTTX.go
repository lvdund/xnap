package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DURXMTTXSupported            int64 = 0
	DURXMTTXNotSupported         int64 = 1
	DURXMTTXSupportedFDMRequired int64 = 2
)

var dURXMTTXConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type DURXMTTX struct {
	Value int64
}

func (ie *DURXMTTX) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dURXMTTXConstraints)
}

func (ie *DURXMTTX) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dURXMTTXConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
