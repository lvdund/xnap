package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bHInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBHInfo)),
}

type BHInfoList struct {
	Value []*BHInfoItem
}

func (ie *BHInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bHInfoListConstraints)
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

func (ie *BHInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bHInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BHInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BHInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
