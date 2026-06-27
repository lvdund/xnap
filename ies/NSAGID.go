package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nSAGIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type NSAGID struct {
	Value int64
}

func (ie *NSAGID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, nSAGIDConstraints)
}

func (ie *NSAGID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(nSAGIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
