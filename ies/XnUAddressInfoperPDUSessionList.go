package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var xnUAddressInfoperPDUSessionListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPDUSessions)),
}

type XnUAddressInfoperPDUSessionList struct {
	Value []*XnUAddressInfoperPDUSessionItem
}

func (ie *XnUAddressInfoperPDUSessionList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(xnUAddressInfoperPDUSessionListConstraints)
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

func (ie *XnUAddressInfoperPDUSessionList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(xnUAddressInfoperPDUSessionListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*XnUAddressInfoperPDUSessionItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(XnUAddressInfoperPDUSessionItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
