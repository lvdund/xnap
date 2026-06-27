package ies

import (
	"github.com/lvdund/asn1go/per"
)

var measObjectContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type MeasObjectContainer struct {
	Value []byte
}

func (ie *MeasObjectContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, measObjectContainerConstraints)
}

func (ie *MeasObjectContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(measObjectContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
