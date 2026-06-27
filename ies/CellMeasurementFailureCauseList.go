package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellMeasurementFailureCauseListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxFailedCellMeasObjects)),
}

type CellMeasurementFailureCauseList struct {
	Value []*CellMeasurementFailureCauseItem
}

func (ie *CellMeasurementFailureCauseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellMeasurementFailureCauseListConstraints)
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

func (ie *CellMeasurementFailureCauseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellMeasurementFailureCauseListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CellMeasurementFailureCauseItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CellMeasurementFailureCauseItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
