package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var parentNodeCellsListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofServingCells)),
}

type ParentNodeCellsList struct {
	Value []*ParentNodeCellsListItem
}

func (ie *ParentNodeCellsList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(parentNodeCellsListConstraints)
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

func (ie *ParentNodeCellsList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(parentNodeCellsListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ParentNodeCellsListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ParentNodeCellsListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
