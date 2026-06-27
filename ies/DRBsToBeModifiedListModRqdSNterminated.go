package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBsToBeModifiedListModRqdSNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DRBsToBeModifiedListModRqdSNterminated struct {
	Value []*DRBsToBeModifiedListModRqdSNterminatedItem
}

func (ie *DRBsToBeModifiedListModRqdSNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBsToBeModifiedListModRqdSNterminatedConstraints)
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

func (ie *DRBsToBeModifiedListModRqdSNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBsToBeModifiedListModRqdSNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DRBsToBeModifiedListModRqdSNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DRBsToBeModifiedListModRqdSNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
