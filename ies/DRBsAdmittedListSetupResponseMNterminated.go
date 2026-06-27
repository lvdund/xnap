package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBsAdmittedListSetupResponseMNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DRBsAdmittedListSetupResponseMNterminated struct {
	Value []*DRBsAdmittedListSetupResponseMNterminatedItem
}

func (ie *DRBsAdmittedListSetupResponseMNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBsAdmittedListSetupResponseMNterminatedConstraints)
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

func (ie *DRBsAdmittedListSetupResponseMNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBsAdmittedListSetupResponseMNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DRBsAdmittedListSetupResponseMNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DRBsAdmittedListSetupResponseMNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
