package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERadioCapabilityIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type UERadioCapabilityID struct {
	Value []byte
}

func (ie *UERadioCapabilityID) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, uERadioCapabilityIDConstraints)
}

func (ie *UERadioCapabilityID) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(uERadioCapabilityIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
