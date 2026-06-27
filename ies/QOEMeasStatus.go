package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QOEMeasStatusOngoing int64 = 0
)

var qOEMeasStatusConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type QOEMeasStatus struct {
	Value int64
}

func (ie *QOEMeasStatus) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qOEMeasStatusConstraints)
}

func (ie *QOEMeasStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qOEMeasStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
