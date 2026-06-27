package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var extendedPacketDelayBudgetConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(65535)),
}

type ExtendedPacketDelayBudget struct {
	Value int64
}

func (ie *ExtendedPacketDelayBudget) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, extendedPacketDelayBudgetConstraints)
}

func (ie *ExtendedPacketDelayBudget) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(extendedPacketDelayBudgetConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
