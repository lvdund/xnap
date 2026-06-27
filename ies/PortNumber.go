package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var portNumberConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(16)),
	Max:        common.Ptr(int64(16)),
}

type PortNumber struct {
	Value per.BitString
}

func (ie *PortNumber) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, portNumberConstraints)
}

func (ie *PortNumber) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(portNumberConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
