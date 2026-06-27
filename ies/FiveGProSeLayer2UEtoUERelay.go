package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeLayer2UEtoUERelayAuthorized    int64 = 0
	FiveGProSeLayer2UEtoUERelayNotAuthorized int64 = 1
)

var fiveGProSeLayer2UEtoUERelayConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeLayer2UEtoUERelay struct {
	Value int64
}

func (ie *FiveGProSeLayer2UEtoUERelay) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeLayer2UEtoUERelayConstraints)
}

func (ie *FiveGProSeLayer2UEtoUERelay) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeLayer2UEtoUERelayConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
