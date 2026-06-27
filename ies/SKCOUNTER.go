package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sKCOUNTERConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(65535)),
}

type SKCOUNTER struct {
	Value int64
}

func (ie *SKCOUNTER) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sKCOUNTERConstraints)
}

func (ie *SKCOUNTER) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sKCOUNTERConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
