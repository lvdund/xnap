package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOMaxnoofCondReconfigConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(8)),
}

type CHOMaxnoofCondReconfig struct {
	Value int64
}

func (ie *CHOMaxnoofCondReconfig) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cHOMaxnoofCondReconfigConstraints)
}

func (ie *CHOMaxnoofCondReconfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cHOMaxnoofCondReconfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
