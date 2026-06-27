package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var associatedQoSFlowInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSQoSFlows)),
}

type AssociatedQoSFlowInfoList struct {
	Value []*AssociatedQoSFlowInfoItem
}

func (ie *AssociatedQoSFlowInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(associatedQoSFlowInfoListConstraints)
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

func (ie *AssociatedQoSFlowInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(associatedQoSFlowInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*AssociatedQoSFlowInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(AssociatedQoSFlowInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
