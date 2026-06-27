package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var wLANNameConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(32)),
}

type WLANName struct {
	Value []byte
}

func (ie *WLANName) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, wLANNameConstraints)
}

func (ie *WLANName) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(wLANNameConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
