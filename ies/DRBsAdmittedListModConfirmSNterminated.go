package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBsAdmittedListModConfirmSNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DRBsAdmittedListModConfirmSNterminated struct {
	Value []*DRBsAdmittedListModConfirmSNterminatedItem
}

func (ie *DRBsAdmittedListModConfirmSNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBsAdmittedListModConfirmSNterminatedConstraints)
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

func (ie *DRBsAdmittedListModConfirmSNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBsAdmittedListModConfirmSNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DRBsAdmittedListModConfirmSNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DRBsAdmittedListModConfirmSNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
