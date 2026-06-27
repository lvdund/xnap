package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var conditionalReconfigListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPSCellCandidates)),
}

type ConditionalReconfigList struct {
	Value []*ConditionalReconfigItem
}

func (ie *ConditionalReconfigList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(conditionalReconfigListConstraints)
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

func (ie *ConditionalReconfigList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(conditionalReconfigListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ConditionalReconfigItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ConditionalReconfigItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
