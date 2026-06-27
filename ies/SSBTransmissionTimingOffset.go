package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBTransmissionTimingOffsetConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(127)),
}

type SSBTransmissionTimingOffset struct {
	Value int64
}

func (ie *SSBTransmissionTimingOffset) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sSBTransmissionTimingOffsetConstraints)
}

func (ie *SSBTransmissionTimingOffset) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sSBTransmissionTimingOffsetConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
