package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERadioCapabilityForPagingOfNRConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type UERadioCapabilityForPagingOfNR struct {
	Value []byte
}

func (ie *UERadioCapabilityForPagingOfNR) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, uERadioCapabilityForPagingOfNRConstraints)
}

func (ie *UERadioCapabilityForPagingOfNR) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(uERadioCapabilityForPagingOfNRConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
