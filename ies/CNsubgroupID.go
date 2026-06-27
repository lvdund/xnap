package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cNsubgroupIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(7)),
}

type CNsubgroupID struct {
	Value int64
}

func (ie *CNsubgroupID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cNsubgroupIDConstraints)
}

func (ie *CNsubgroupID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cNsubgroupIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
