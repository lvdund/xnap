package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionIDConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type PDUSessionID struct {
	Value int64
}

func (ie *PDUSessionID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, pDUSessionIDConstraints)
}

func (ie *PDUSessionID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(pDUSessionIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
