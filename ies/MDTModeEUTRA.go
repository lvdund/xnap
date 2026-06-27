package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mDTModeEUTRAConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type MDTModeEUTRA struct {
	Value []byte
}

func (ie *MDTModeEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, mDTModeEUTRAConstraints)
}

func (ie *MDTModeEUTRA) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(mDTModeEUTRAConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
