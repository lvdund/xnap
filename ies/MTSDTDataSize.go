package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mTSDTDataSizeConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(96000)),
}

type MTSDTDataSize struct {
	Value int64
}

func (ie *MTSDTDataSize) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, mTSDTDataSizeConstraints)
}

func (ie *MTSDTDataSize) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(mTSDTDataSizeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
