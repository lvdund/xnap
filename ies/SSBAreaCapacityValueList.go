package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBAreaCapacityValueListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
}

type SSBAreaCapacityValueList struct {
	Value []*SSBAreaCapacityValueListItem
}

func (ie *SSBAreaCapacityValueList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBAreaCapacityValueListConstraints)
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

func (ie *SSBAreaCapacityValueList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBAreaCapacityValueListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SSBAreaCapacityValueListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SSBAreaCapacityValueListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
