package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var boundaryNodeCellsListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofServedCellsIAB)),
}

type BoundaryNodeCellsList struct {
	Value []*BoundaryNodeCellsListItem
}

func (ie *BoundaryNodeCellsList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(boundaryNodeCellsListConstraints)
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

func (ie *BoundaryNodeCellsList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(boundaryNodeCellsListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BoundaryNodeCellsListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BoundaryNodeCellsListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
