package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var maskedIMEISVConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(64)),
	Max:        common.Ptr(int64(64)),
}

type MaskedIMEISV struct {
	Value per.BitString
}

func (ie *MaskedIMEISV) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, maskedIMEISVConstraints)
}

func (ie *MaskedIMEISV) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(maskedIMEISVConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
