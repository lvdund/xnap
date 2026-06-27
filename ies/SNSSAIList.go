package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNSSAIListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSliceItems)),
}

type SNSSAIList struct {
	Value []*SNSSAIItem
}

func (ie *SNSSAIList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNSSAIListConstraints)
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

func (ie *SNSSAIList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNSSAIListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SNSSAIItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SNSSAIItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
