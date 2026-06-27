package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var m7periodConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(60)),
}

type M7period struct {
	Value int64
}

func (ie *M7period) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, m7periodConstraints)
}

func (ie *M7period) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(m7periodConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
