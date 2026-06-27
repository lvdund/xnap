package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NoPDUSessionIndicationTrue int64 = 0
)

var noPDUSessionIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type NoPDUSessionIndication struct {
	Value int64
}

func (ie *NoPDUSessionIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, noPDUSessionIndicationConstraints)
}

func (ie *NoPDUSessionIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(noPDUSessionIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
