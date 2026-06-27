package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RRCConfigIndicationFullConfig  int64 = 0
	RRCConfigIndicationDeltaConfig int64 = 1
)

var rRCConfigIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RRCConfigIndication struct {
	Value int64
}

func (ie *RRCConfigIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rRCConfigIndicationConstraints)
}

func (ie *RRCConfigIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rRCConfigIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
