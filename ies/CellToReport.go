package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellToReportConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
}

type CellToReport struct {
	Value []*CellToReportItem
}

func (ie *CellToReport) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellToReportConstraints)
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

func (ie *CellToReport) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellToReportConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CellToReportItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CellToReportItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
