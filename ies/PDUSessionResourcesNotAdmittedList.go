package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionResourcesNotAdmittedListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionResourcesNotAdmittedList struct {
	Value []*PDUSessionResourcesNotAdmittedItem
}

func (ie *PDUSessionResourcesNotAdmittedList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionResourcesNotAdmittedListConstraints)
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

func (ie *PDUSessionResourcesNotAdmittedList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionResourcesNotAdmittedListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionResourcesNotAdmittedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionResourcesNotAdmittedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
