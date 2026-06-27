package ies

import (
	"github.com/lvdund/asn1go/per"
)

var f1CTrafficContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type F1CTrafficContainer struct {
	Value []byte
}

func (ie *F1CTrafficContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, f1CTrafficContainerConstraints)
}

func (ie *F1CTrafficContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(f1CTrafficContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
