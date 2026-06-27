package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var additionalULNGUTNLatUPFListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMultiConnectivityMinusOne)),
}

type AdditionalULNGUTNLatUPFList struct {
	Value []*AdditionalULNGUTNLatUPFItem
}

func (ie *AdditionalULNGUTNLatUPFList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(additionalULNGUTNLatUPFListConstraints)
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

func (ie *AdditionalULNGUTNLatUPFList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(additionalULNGUTNLatUPFListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*AdditionalULNGUTNLatUPFItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(AdditionalULNGUTNLatUPFItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
