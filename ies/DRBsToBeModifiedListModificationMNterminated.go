package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBsToBeModifiedListModificationMNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DRBsToBeModifiedListModificationMNterminated struct {
	Value []*DRBsToBeModifiedListModificationMNterminatedItem
}

func (ie *DRBsToBeModifiedListModificationMNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBsToBeModifiedListModificationMNterminatedConstraints)
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

func (ie *DRBsToBeModifiedListModificationMNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBsToBeModifiedListModificationMNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DRBsToBeModifiedListModificationMNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DRBsToBeModifiedListModificationMNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
