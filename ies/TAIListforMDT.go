package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tAIListforMDTConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTAforMDT)),
}

type TAIListforMDT struct {
	Value []*TAIforMDTItem
}

func (ie *TAIListforMDT) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tAIListforMDTConstraints)
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

func (ie *TAIListforMDT) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tAIListforMDTConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TAIforMDTItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TAIforMDTItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
