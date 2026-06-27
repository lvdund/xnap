package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tNLAToRemoveListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTNLAssociations)),
}

type TNLAToRemoveList struct {
	Value []*TNLAToRemoveItem
}

func (ie *TNLAToRemoveList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tNLAToRemoveListConstraints)
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

func (ie *TNLAToRemoveList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tNLAToRemoveListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TNLAToRemoveItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TNLAToRemoveItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
