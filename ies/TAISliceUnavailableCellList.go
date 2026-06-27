package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tAISliceUnavailableCellListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofExtSliceItems)),
}

type TAISliceUnavailableCellList struct {
	Value []*TAISliceUnavailableCellItem
}

func (ie *TAISliceUnavailableCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tAISliceUnavailableCellListConstraints)
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

func (ie *TAISliceUnavailableCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tAISliceUnavailableCellListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TAISliceUnavailableCellItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TAISliceUnavailableCellItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
