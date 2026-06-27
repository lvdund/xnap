package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSQoSFlowsToAddListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSQoSFlows)),
}

type MBSQoSFlowsToAddList struct {
	Value []*MBSQoSFlowsToAddItem
}

func (ie *MBSQoSFlowsToAddList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSQoSFlowsToAddListConstraints)
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

func (ie *MBSQoSFlowsToAddList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSQoSFlowsToAddListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSQoSFlowsToAddItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSQoSFlowsToAddItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
