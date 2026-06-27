package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var resetResponsePartialReleaseListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofUEContexts)),
}

type ResetResponsePartialReleaseList struct {
	Value []*ResetResponsePartialReleaseItem
}

func (ie *ResetResponsePartialReleaseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(resetResponsePartialReleaseListConstraints)
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

func (ie *ResetResponsePartialReleaseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(resetResponsePartialReleaseListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ResetResponsePartialReleaseItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ResetResponsePartialReleaseItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
