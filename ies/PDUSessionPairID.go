package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionPairIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type PDUSessionPairID struct {
	Value int64
}

func (ie *PDUSessionPairID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, pDUSessionPairIDConstraints)
}

func (ie *PDUSessionPairID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(pDUSessionPairIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
