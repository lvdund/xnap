package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSDataForwardingResponseInfofromTargetConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMRBs)),
}

type MBSDataForwardingResponseInfofromTarget struct {
	Value []*MBSDataForwardingResponseInfofromTargetItem
}

func (ie *MBSDataForwardingResponseInfofromTarget) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSDataForwardingResponseInfofromTargetConstraints)
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

func (ie *MBSDataForwardingResponseInfofromTarget) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSDataForwardingResponseInfofromTargetConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSDataForwardingResponseInfofromTargetItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSDataForwardingResponseInfofromTargetItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
