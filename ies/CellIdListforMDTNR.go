package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellIdListforMDTNRConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellIDforMDT)),
}

type CellIdListforMDTNR struct {
	Value []*NRCGI
}

func (ie *CellIdListforMDTNR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellIdListforMDTNRConstraints)
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

func (ie *CellIdListforMDTNR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellIdListforMDTNRConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NRCGI, n)
	for i := range ie.Value {
		ie.Value[i] = new(NRCGI)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
