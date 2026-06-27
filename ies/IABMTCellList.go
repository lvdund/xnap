package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var iABMTCellListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofServingCells)),
}

type IABMTCellList struct {
	Value []*IABMTCellListItem
}

func (ie *IABMTCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(iABMTCellListConstraints)
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

func (ie *IABMTCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(iABMTCellListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*IABMTCellListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(IABMTCellListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
