package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qOEReferenceConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(6)),
	Max:        common.Ptr(int64(6)),
}

type QOEReference struct {
	Value []byte
}

func (ie *QOEReference) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, qOEReferenceConstraints)
}

func (ie *QOEReference) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(qOEReferenceConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
