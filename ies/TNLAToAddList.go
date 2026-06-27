package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tNLAToAddListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTNLAssociations)),
}

type TNLAToAddList struct {
	Value []*TNLAToAddItem
}

func (ie *TNLAToAddList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tNLAToAddListConstraints)
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

func (ie *TNLAToAddList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tNLAToAddListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TNLAToAddItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TNLAToAddItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
