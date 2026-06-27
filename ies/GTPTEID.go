package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var gTPTEIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(4)),
	Max:        common.Ptr(int64(4)),
}

type GTPTEID struct {
	Value []byte
}

func (ie *GTPTEID) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, gTPTEIDConstraints)
}

func (ie *GTPTEID) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(gTPTEIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
