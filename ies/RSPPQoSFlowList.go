package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var rSPPQoSFlowListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofRSPPQoSFlows)),
}

type RSPPQoSFlowList struct {
	Value []*RSPPQoSFlowItem
}

func (ie *RSPPQoSFlowList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rSPPQoSFlowListConstraints)
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

func (ie *RSPPQoSFlowList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rSPPQoSFlowListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*RSPPQoSFlowItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(RSPPQoSFlowItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
