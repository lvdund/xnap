package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDUSetbasedHandlingIndicatorSupported int64 = 0
)

var pDUSetbasedHandlingIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PDUSetbasedHandlingIndicator struct {
	Value int64
}

func (ie *PDUSetbasedHandlingIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDUSetbasedHandlingIndicatorConstraints)
}

func (ie *PDUSetbasedHandlingIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDUSetbasedHandlingIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
