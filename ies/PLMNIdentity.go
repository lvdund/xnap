package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pLMNIdentityConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(3)),
	Max:        common.Ptr(int64(3)),
}

type PLMNIdentity struct {
	Value []byte
}

func (ie *PLMNIdentity) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, pLMNIdentityConstraints)
}

func (ie *PLMNIdentity) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(pLMNIdentityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
