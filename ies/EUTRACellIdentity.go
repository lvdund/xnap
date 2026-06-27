package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eUTRACellIdentityConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(28)),
	Max:        common.Ptr(int64(28)),
}

type EUTRACellIdentity struct {
	Value per.BitString
}

func (ie *EUTRACellIdentity) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, eUTRACellIdentityConstraints)
}

func (ie *EUTRACellIdentity) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(eUTRACellIdentityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
