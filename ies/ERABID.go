package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eRABIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(15)),
}

type ERABID struct {
	Value int64
}

func (ie *ERABID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, eRABIDConstraints)
}

func (ie *ERABID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(eRABIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
