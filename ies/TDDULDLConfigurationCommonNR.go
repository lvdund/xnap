package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tDDULDLConfigurationCommonNRConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type TDDULDLConfigurationCommonNR struct {
	Value []byte
}

func (ie *TDDULDLConfigurationCommonNR) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, tDDULDLConfigurationCommonNRConstraints)
}

func (ie *TDDULDLConfigurationCommonNR) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(tDDULDLConfigurationCommonNRConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
