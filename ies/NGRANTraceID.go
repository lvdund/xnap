package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nGRANTraceIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(8)),
	Max:        common.Ptr(int64(8)),
}

type NGRANTraceID struct {
	Value []byte
}

func (ie *NGRANTraceID) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, nGRANTraceIDConstraints)
}

func (ie *NGRANTraceID) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(nGRANTraceIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
