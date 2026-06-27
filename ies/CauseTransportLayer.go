package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CauseTransportLayerTransportResourceUnavailable int64 = 0
	CauseTransportLayerUnspecified                  int64 = 1
)

var causeTransportLayerConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type CauseTransportLayer struct {
	Value int64
}

func (ie *CauseTransportLayer) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, causeTransportLayerConstraints)
}

func (ie *CauseTransportLayer) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(causeTransportLayerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
