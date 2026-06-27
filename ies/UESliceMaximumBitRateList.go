package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uESliceMaximumBitRateListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSMBR)),
}

type UESliceMaximumBitRateList struct {
	Value []*UESliceMaximumBitRateItem
}

func (ie *UESliceMaximumBitRateList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uESliceMaximumBitRateListConstraints)
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

func (ie *UESliceMaximumBitRateList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uESliceMaximumBitRateListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*UESliceMaximumBitRateItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(UESliceMaximumBitRateItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
