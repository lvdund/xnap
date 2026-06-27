package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var servedCellsNRConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
}

type ServedCellsNR struct {
	Value []*ServedCellsNRItem
}

func (ie *ServedCellsNR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(servedCellsNRConstraints)
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

func (ie *ServedCellsNR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(servedCellsNRConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ServedCellsNRItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ServedCellsNRItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
