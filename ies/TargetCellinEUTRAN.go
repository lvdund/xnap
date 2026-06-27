package ies

import (
	"github.com/lvdund/asn1go/per"
)

var targetCellinEUTRANConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type TargetCellinEUTRAN struct {
	Value []byte
}

func (ie *TargetCellinEUTRAN) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, targetCellinEUTRANConstraints)
}

func (ie *TargetCellinEUTRAN) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(targetCellinEUTRANConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
