package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var trafficIndexConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(1024)),
}

type TrafficIndex struct {
	Value int64
}

func (ie *TrafficIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, trafficIndexConstraints)
}

func (ie *TrafficIndex) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(trafficIndexConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
