package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rACHConfigCommonConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type RACHConfigCommon struct {
	Value []byte
}

func (ie *RACHConfigCommon) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, rACHConfigCommonConstraints)
}

func (ie *RACHConfigCommon) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(rACHConfigCommonConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
