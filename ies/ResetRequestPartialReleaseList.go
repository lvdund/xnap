package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var resetRequestPartialReleaseListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofUEContexts)),
}

type ResetRequestPartialReleaseList struct {
	Value []*ResetRequestPartialReleaseItem
}

func (ie *ResetRequestPartialReleaseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(resetRequestPartialReleaseListConstraints)
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

func (ie *ResetRequestPartialReleaseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(resetRequestPartialReleaseListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ResetRequestPartialReleaseItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ResetRequestPartialReleaseItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
