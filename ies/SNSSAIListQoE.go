package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNSSAIListQoEConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSNSSAIforQMC)),
}

type SNSSAIListQoE struct {
	Value []*SNSSAI
}

func (ie *SNSSAIListQoE) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNSSAIListQoEConstraints)
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

func (ie *SNSSAIListQoE) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNSSAIListQoEConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SNSSAI, n)
	for i := range ie.Value {
		ie.Value[i] = new(SNSSAI)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
