package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var fiveGProSePC5QoSFlowListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPC5QoSFlows)),
}

type FiveGProSePC5QoSFlowList struct {
	Value []*FiveGProSePC5QoSFlowItem
}

func (ie *FiveGProSePC5QoSFlowList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(fiveGProSePC5QoSFlowListConstraints)
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

func (ie *FiveGProSePC5QoSFlowList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(fiveGProSePC5QoSFlowListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*FiveGProSePC5QoSFlowItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(FiveGProSePC5QoSFlowItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
