package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionResourcesActivityNotifyListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type PDUSessionResourcesActivityNotifyList struct {
	Value []*PDUSessionResourcesActivityNotifyItem
}

func (ie *PDUSessionResourcesActivityNotifyList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDUSessionResourcesActivityNotifyListConstraints)
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

func (ie *PDUSessionResourcesActivityNotifyList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDUSessionResourcesActivityNotifyListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PDUSessionResourcesActivityNotifyItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PDUSessionResourcesActivityNotifyItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
