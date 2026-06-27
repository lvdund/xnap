package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var packetLossRateConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(1000)),
}

type PacketLossRate struct {
	Value int64
}

func (ie *PacketLossRate) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, packetLossRateConstraints)
}

func (ie *PacketLossRate) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(packetLossRateConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
