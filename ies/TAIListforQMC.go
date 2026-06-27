package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tAIListforQMCConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTAforQMC)),
}

type TAIListforQMC struct {
	Value []*TAIItem
}

func (ie *TAIListforQMC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tAIListforQMCConstraints)
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

func (ie *TAIListforQMC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tAIListforQMCConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TAIItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TAIItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
