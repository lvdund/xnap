package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellIdListforMDTEUTRAConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellIDforMDT)),
}

type CellIdListforMDTEUTRA struct {
	Value []*EUTRACGI
}

func (ie *CellIdListforMDTEUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellIdListforMDTEUTRAConstraints)
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

func (ie *CellIdListforMDTEUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellIdListforMDTEUTRAConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*EUTRACGI, n)
	for i := range ie.Value {
		ie.Value[i] = new(EUTRACGI)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
