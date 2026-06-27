package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var requestReferenceIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(64)),
}

type RequestReferenceID struct {
	Value int64
}

func (ie *RequestReferenceID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, requestReferenceIDConstraints)
}

func (ie *RequestReferenceID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(requestReferenceIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
