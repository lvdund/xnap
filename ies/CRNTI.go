package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cRNTIConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(16)),
	Max:        common.Ptr(int64(16)),
}

type CRNTI struct {
	Value per.BitString
}

func (ie *CRNTI) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, cRNTIConstraints)
}

func (ie *CRNTI) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(cRNTIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
