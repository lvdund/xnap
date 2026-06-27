package ies

import (
	"github.com/lvdund/asn1go/per"
)

var routingIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type RoutingID struct {
	Value []byte
}

func (ie *RoutingID) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, routingIDConstraints)
}

func (ie *RoutingID) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(routingIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
