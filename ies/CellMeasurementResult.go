package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellMeasurementResultConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
}

type CellMeasurementResult struct {
	Value []*CellMeasurementResultItem
}

func (ie *CellMeasurementResult) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellMeasurementResultConstraints)
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

func (ie *CellMeasurementResult) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellMeasurementResultConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CellMeasurementResultItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CellMeasurementResultItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
