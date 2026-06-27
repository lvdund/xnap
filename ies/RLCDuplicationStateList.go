package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var rLCDuplicationStateListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofRLCDuplicationstate)),
}

type RLCDuplicationStateList struct {
	Value []*RLCDuplicationStateItem
}

func (ie *RLCDuplicationStateList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rLCDuplicationStateListConstraints)
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

func (ie *RLCDuplicationStateList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rLCDuplicationStateListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*RLCDuplicationStateItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(RLCDuplicationStateItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
