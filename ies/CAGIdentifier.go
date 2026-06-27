package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cAGIdentifierConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(32)),
	Max:        common.Ptr(int64(32)),
}

type CAGIdentifier struct {
	Value per.BitString
}

func (ie *CAGIdentifier) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, cAGIdentifierConstraints)
}

func (ie *CAGIdentifier) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(cAGIdentifierConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
