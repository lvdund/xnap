package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CSIRSTransmissionIndicationActivated   int64 = 0
	CSIRSTransmissionIndicationDeactivated int64 = 1
)

var cSIRSTransmissionIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type CSIRSTransmissionIndication struct {
	Value int64
}

func (ie *CSIRSTransmissionIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cSIRSTransmissionIndicationConstraints)
}

func (ie *CSIRSTransmissionIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cSIRSTransmissionIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
