package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var packetDelayBudgetConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(1023)),
}

type PacketDelayBudget struct {
	Value int64
}

func (ie *PacketDelayBudget) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, packetDelayBudgetConstraints)
}

func (ie *PacketDelayBudget) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(packetDelayBudgetConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
