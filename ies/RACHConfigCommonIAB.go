package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rACHConfigCommonIABConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type RACHConfigCommonIAB struct {
	Value []byte
}

func (ie *RACHConfigCommonIAB) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, rACHConfigCommonIABConstraints)
}

func (ie *RACHConfigCommonIAB) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(rACHConfigCommonIABConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
