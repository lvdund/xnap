package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pC5QoSFlowListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPC5QoSFlows)),
}

type PC5QoSFlowList struct {
	Value []*PC5QoSFlowItem
}

func (ie *PC5QoSFlowList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pC5QoSFlowListConstraints)
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

func (ie *PC5QoSFlowList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pC5QoSFlowListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PC5QoSFlowItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PC5QoSFlowItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
