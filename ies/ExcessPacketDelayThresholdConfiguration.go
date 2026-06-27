package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var excessPacketDelayThresholdConfigurationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofThresholdsForExcessPacketDelay)),
}

type ExcessPacketDelayThresholdConfiguration struct {
	Value []*ExcessPacketDelayThresholdItem
}

func (ie *ExcessPacketDelayThresholdConfiguration) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(excessPacketDelayThresholdConfigurationConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ExcessPacketDelayThresholdConfiguration) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(excessPacketDelayThresholdConfigurationConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ExcessPacketDelayThresholdItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ExcessPacketDelayThresholdItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
