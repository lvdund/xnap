package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionsToBeModifiedSNModRequestListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionsToBeModifiedSNModRequestList struct {
	Value []*PDUSessionsToBeModifiedSNModRequestItem
}

func (ie *PDUSessionsToBeModifiedSNModRequestList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionsToBeModifiedSNModRequestListConstraints)
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

func (ie *PDUSessionsToBeModifiedSNModRequestList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionsToBeModifiedSNModRequestListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionsToBeModifiedSNModRequestItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionsToBeModifiedSNModRequestItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
