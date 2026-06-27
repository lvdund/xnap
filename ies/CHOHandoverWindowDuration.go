package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOHandoverWindowDurationConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(6000)),
}

type CHOHandoverWindowDuration struct {
	Value int64
}

func (ie *CHOHandoverWindowDuration) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cHOHandoverWindowDurationConstraints)
}

func (ie *CHOHandoverWindowDuration) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cHOHandoverWindowDurationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
