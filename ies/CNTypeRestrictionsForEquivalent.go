package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cNTypeRestrictionsForEquivalentConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofEPLMNs)),
}

type CNTypeRestrictionsForEquivalent struct {
	Value []*CNTypeRestrictionsForEquivalentItem
}

func (ie *CNTypeRestrictionsForEquivalent) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cNTypeRestrictionsForEquivalentConstraints)
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

func (ie *CNTypeRestrictionsForEquivalent) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cNTypeRestrictionsForEquivalentConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CNTypeRestrictionsForEquivalentItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CNTypeRestrictionsForEquivalentItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
