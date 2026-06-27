package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bearersSubjectToCounterCheckListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type BearersSubjectToCounterCheckList struct {
	Value []*BearersSubjectToCounterCheckItem
}

func (ie *BearersSubjectToCounterCheckList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bearersSubjectToCounterCheckListConstraints)
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

func (ie *BearersSubjectToCounterCheckList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bearersSubjectToCounterCheckListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BearersSubjectToCounterCheckItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BearersSubjectToCounterCheckItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
