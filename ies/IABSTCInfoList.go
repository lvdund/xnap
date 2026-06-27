package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var iABSTCInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofIABSTCInfo)),
}

type IABSTCInfoList struct {
	Value []*IABSTCInfoItem
}

func (ie *IABSTCInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(iABSTCInfoListConstraints)
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

func (ie *IABSTCInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(iABSTCInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*IABSTCInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(IABSTCInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
