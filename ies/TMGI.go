package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tMGIConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(6)),
	Max:        common.Ptr(int64(6)),
}

type TMGI struct {
	Value []byte
}

func (ie *TMGI) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, tMGIConstraints)
}

func (ie *TMGI) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(tMGIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
