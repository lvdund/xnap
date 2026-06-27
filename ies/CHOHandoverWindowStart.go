package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOHandoverWindowStartConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(549755813887)),
}

type CHOHandoverWindowStart struct {
	Value int64
}

func (ie *CHOHandoverWindowStart) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cHOHandoverWindowStartConstraints)
}

func (ie *CHOHandoverWindowStart) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cHOHandoverWindowStartConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
