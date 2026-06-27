package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var additionalListofPDUSessionResourceChangeConfirmInfoSNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTargetSNsMinusOne)),
}

type AdditionalListofPDUSessionResourceChangeConfirmInfoSNterminated struct {
	Value []*AdditionalListofPDUSessionResourceChangeConfirmInfoSNterminatedItem
}

func (ie *AdditionalListofPDUSessionResourceChangeConfirmInfoSNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(additionalListofPDUSessionResourceChangeConfirmInfoSNterminatedConstraints)
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

func (ie *AdditionalListofPDUSessionResourceChangeConfirmInfoSNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(additionalListofPDUSessionResourceChangeConfirmInfoSNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*AdditionalListofPDUSessionResourceChangeConfirmInfoSNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(AdditionalListofPDUSessionResourceChangeConfirmInfoSNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
