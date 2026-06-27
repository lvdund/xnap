package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var broadcastEUTRAPLMNsConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofEUTRABPLMNs)),
}

type BroadcastEUTRAPLMNs struct {
	Value []*PLMNIdentity
}

func (ie *BroadcastEUTRAPLMNs) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(broadcastEUTRAPLMNsConstraints)
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

func (ie *BroadcastEUTRAPLMNs) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(broadcastEUTRAPLMNsConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PLMNIdentity, n)
	for i := range ie.Value {
		ie.Value[i] = new(PLMNIdentity)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
