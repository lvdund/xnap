package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var a2XPC5QoSFlowListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPC5QoSFlows)),
}

type A2XPC5QoSFlowList struct {
	Value []*A2XPC5QoSFlowItem
}

func (ie *A2XPC5QoSFlowList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(a2XPC5QoSFlowListConstraints)
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

func (ie *A2XPC5QoSFlowList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(a2XPC5QoSFlowListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*A2XPC5QoSFlowItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(A2XPC5QoSFlowItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
