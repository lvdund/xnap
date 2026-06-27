package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uEAppLayerMeasInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofUEAppLayerMeas)),
}

type UEAppLayerMeasInfoList struct {
	Value []*UEAppLayerMeasInfoItem
}

func (ie *UEAppLayerMeasInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uEAppLayerMeasInfoListConstraints)
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

func (ie *UEAppLayerMeasInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uEAppLayerMeasInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*UEAppLayerMeasInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(UEAppLayerMeasInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
