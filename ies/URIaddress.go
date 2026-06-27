package ies

import (
	"github.com/lvdund/asn1go/per"
)

type URIaddress struct {
	Value string
}

func (ie *URIaddress) Encode(e *per.Encoder) error {
	return e.EncodeRestrictedString(ie.Value, per.CharacterStringConstraints{TypeName: "VisibleString"})
}

func (ie *URIaddress) Decode(d *per.Decoder) error {
	val, err := d.DecodeRestrictedString(per.CharacterStringConstraints{TypeName: "VisibleString"})
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
