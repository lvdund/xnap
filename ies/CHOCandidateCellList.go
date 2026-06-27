package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOCandidateCellListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinCHO)),
}

type CHOCandidateCellList struct {
	Value []*CHOCandidateCellItem
}

func (ie *CHOCandidateCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cHOCandidateCellListConstraints)
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

func (ie *CHOCandidateCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cHOCandidateCellListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CHOCandidateCellItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CHOCandidateCellItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
