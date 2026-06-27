package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionListWithDataForwardingRequestConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionListWithDataForwardingRequest struct {
	Value []*PDUSessionListWithDataForwardingRequestItem
}

func (ie *PDUSessionListWithDataForwardingRequest) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionListWithDataForwardingRequestConstraints)
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

func (ie *PDUSessionListWithDataForwardingRequest) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionListWithDataForwardingRequestConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionListWithDataForwardingRequestItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionListWithDataForwardingRequestItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
