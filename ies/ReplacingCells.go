package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var replacingCellsConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(0)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
}

type ReplacingCells struct {
	Value []*ReplacingCellsItem
}

func (ie *ReplacingCells) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(replacingCellsConstraints)
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

func (ie *ReplacingCells) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(replacingCellsConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ReplacingCellsItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ReplacingCellsItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
