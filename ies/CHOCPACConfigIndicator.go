package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CHOCPACConfigIndicatorChoOnlyNotPrepared int64 = 0
)

var cHOCPACConfigIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CHOCPACConfigIndicator struct {
	Value int64
}

func (ie *CHOCPACConfigIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cHOCPACConfigIndicatorConstraints)
}

func (ie *CHOCPACConfigIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cHOCPACConfigIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
