package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBFreqInfoConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(common.MaxNRARFCN)),
}

type SSBFreqInfo struct {
	Value int64
}

func (ie *SSBFreqInfo) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sSBFreqInfoConstraints)
}

func (ie *SSBFreqInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sSBFreqInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
