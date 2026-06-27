package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionAdmittedModSNModConfirmConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionAdmittedModSNModConfirm struct {
	Value []*PDUSessionAdmittedModSNModConfirmItem
}

func (ie *PDUSessionAdmittedModSNModConfirm) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionAdmittedModSNModConfirmConstraints)
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

func (ie *PDUSessionAdmittedModSNModConfirm) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionAdmittedModSNModConfirmConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionAdmittedModSNModConfirmItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionAdmittedModSNModConfirmItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
