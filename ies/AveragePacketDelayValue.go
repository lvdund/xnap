package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var averagePacketDelayValueConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(10000)),
}

type AveragePacketDelayValue struct {
	Value int64
}

func (ie *AveragePacketDelayValue) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, averagePacketDelayValueConstraints)
}

func (ie *AveragePacketDelayValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(averagePacketDelayValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
