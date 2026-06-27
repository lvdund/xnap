package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sRSConfigurationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type SRSConfiguration struct {
	Value []byte
}

func (ie *SRSConfiguration) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, sRSConfigurationConstraints)
}

func (ie *SRSConfiguration) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(sRSConfigurationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
