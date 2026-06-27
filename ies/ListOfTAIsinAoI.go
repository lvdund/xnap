package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var listOfTAIsinAoIConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTAIsinAoI)),
}

type ListOfTAIsinAoI struct {
	Value []*TAIsinAoIItem
}

func (ie *ListOfTAIsinAoI) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(listOfTAIsinAoIConstraints)
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

func (ie *ListOfTAIsinAoI) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(listOfTAIsinAoIConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TAIsinAoIItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TAIsinAoIItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
