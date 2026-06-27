package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bHInfoIndexConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(common.MaxnoofBHInfo)),
}

type BHInfoIndex struct {
	Value int64
}

func (ie *BHInfoIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, bHInfoIndexConstraints)
}

func (ie *BHInfoIndex) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(bHInfoIndexConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
