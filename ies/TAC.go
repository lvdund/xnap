package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tACConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(3)),
	Max:        common.Ptr(int64(3)),
}

type TAC struct {
	Value []byte
}

func (ie *TAC) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, tACConstraints)
}

func (ie *TAC) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(tACConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
