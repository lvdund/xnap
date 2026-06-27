package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionAdmittedToBeModifiedSNModResponseConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionAdmittedToBeModifiedSNModResponse struct {
	Value []*PDUSessionAdmittedToBeModifiedSNModResponseItem
}

func (ie *PDUSessionAdmittedToBeModifiedSNModResponse) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionAdmittedToBeModifiedSNModResponseConstraints)
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

func (ie *PDUSessionAdmittedToBeModifiedSNModResponse) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionAdmittedToBeModifiedSNModResponseConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionAdmittedToBeModifiedSNModResponseItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionAdmittedToBeModifiedSNModResponseItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
