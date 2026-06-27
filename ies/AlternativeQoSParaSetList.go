package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var alternativeQoSParaSetListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSParaSets)),
}

type AlternativeQoSParaSetList struct {
	Value []*AlternativeQoSParaSetItem
}

func (ie *AlternativeQoSParaSetList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(alternativeQoSParaSetListConstraints)
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

func (ie *AlternativeQoSParaSetList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(alternativeQoSParaSetListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*AlternativeQoSParaSetItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(AlternativeQoSParaSetItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
