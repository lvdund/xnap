package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRCellIdentityConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(36)),
	Max:        common.Ptr(int64(36)),
}

type NRCellIdentity struct {
	Value per.BitString
}

func (ie *NRCellIdentity) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, nRCellIdentityConstraints)
}

func (ie *NRCellIdentity) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(nRCellIdentityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
