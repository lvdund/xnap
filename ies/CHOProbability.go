package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOProbabilityConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(100)),
}

type CHOProbability struct {
	Value int64
}

func (ie *CHOProbability) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, cHOProbabilityConstraints)
}

func (ie *CHOProbability) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(cHOProbabilityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
