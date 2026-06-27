package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dAPSResponseInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DAPSResponseInfoList struct {
	Value []*DAPSResponseInfoItem
}

func (ie *DAPSResponseInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dAPSResponseInfoListConstraints)
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

func (ie *DAPSResponseInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dAPSResponseInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DAPSResponseInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DAPSResponseInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
