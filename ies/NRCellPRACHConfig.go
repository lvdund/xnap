package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRCellPRACHConfigConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type NRCellPRACHConfig struct {
	Value []byte
}

func (ie *NRCellPRACHConfig) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, nRCellPRACHConfigConstraints)
}

func (ie *NRCellPRACHConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(nRCellPRACHConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
