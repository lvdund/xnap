package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionSNChangeConfirmListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionSNChangeConfirmList struct {
	Value []*PDUSessionSNChangeConfirmItem
}

func (ie *PDUSessionSNChangeConfirmList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionSNChangeConfirmListConstraints)
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

func (ie *PDUSessionSNChangeConfirmList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionSNChangeConfirmListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionSNChangeConfirmItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionSNChangeConfirmItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
