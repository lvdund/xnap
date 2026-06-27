package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSServiceAreaTAIListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTAIforMBS)),
}

type MBSServiceAreaTAIList struct {
	Value []*MBSServiceAreaTAIItem
}

func (ie *MBSServiceAreaTAIList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSServiceAreaTAIListConstraints)
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

func (ie *MBSServiceAreaTAIList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSServiceAreaTAIListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSServiceAreaTAIItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSServiceAreaTAIItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
