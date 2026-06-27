package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var forbiddenAreaListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPLMNs)),
}

type ForbiddenAreaList struct {
	Value []*ForbiddenAreaItem
}

func (ie *ForbiddenAreaList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(forbiddenAreaListConstraints)
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

func (ie *ForbiddenAreaList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(forbiddenAreaListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ForbiddenAreaItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ForbiddenAreaItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
