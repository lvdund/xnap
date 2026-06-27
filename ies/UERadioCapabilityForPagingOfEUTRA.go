package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERadioCapabilityForPagingOfEUTRAConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type UERadioCapabilityForPagingOfEUTRA struct {
	Value []byte
}

func (ie *UERadioCapabilityForPagingOfEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, uERadioCapabilityForPagingOfEUTRAConstraints)
}

func (ie *UERadioCapabilityForPagingOfEUTRA) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(uERadioCapabilityForPagingOfEUTRAConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
