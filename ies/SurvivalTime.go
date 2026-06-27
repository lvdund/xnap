package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var survivalTimeConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(1920000)),
}

type SurvivalTime struct {
	Value int64
}

func (ie *SurvivalTime) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, survivalTimeConstraints)
}

func (ie *SurvivalTime) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(survivalTimeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
