package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOExecutionConditionListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCHOexecutioncond)),
}

type CHOExecutionConditionList struct {
	Value []*CHOExecutionConditionItem
}

func (ie *CHOExecutionConditionList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cHOExecutionConditionListConstraints)
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

func (ie *CHOExecutionConditionList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cHOExecutionConditionListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CHOExecutionConditionItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CHOExecutionConditionItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
