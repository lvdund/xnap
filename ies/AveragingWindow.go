package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var averagingWindowConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(4095)),
}

type AveragingWindow struct {
	Value int64
}

func (ie *AveragingWindow) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, averagingWindowConstraints)
}

func (ie *AveragingWindow) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(averagingWindowConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
