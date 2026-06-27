package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionsToBeAddedSNModRequestListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionsToBeAddedSNModRequestList struct {
	Value []*PDUSessionsToBeAddedSNModRequestItem
}

func (ie *PDUSessionsToBeAddedSNModRequestList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionsToBeAddedSNModRequestListConstraints)
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

func (ie *PDUSessionsToBeAddedSNModRequestList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionsToBeAddedSNModRequestListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionsToBeAddedSNModRequestItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionsToBeAddedSNModRequestItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
