package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var rANACConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type RANAC struct {
	Value int64
}

func (ie *RANAC) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, rANACConstraints)
}

func (ie *RANAC) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(rANACConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
