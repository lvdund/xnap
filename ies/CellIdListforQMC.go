package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellIdListforQMCConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellIDforQMC)),
}

type CellIdListforQMC struct {
	Value []*GlobalNGRANCellID
}

func (ie *CellIdListforQMC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellIdListforQMCConstraints)
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

func (ie *CellIdListforQMC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellIdListforQMCConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*GlobalNGRANCellID, n)
	for i := range ie.Value {
		ie.Value[i] = new(GlobalNGRANCellID)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
