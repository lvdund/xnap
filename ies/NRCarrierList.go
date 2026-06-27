package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRCarrierListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofNRSCSs)),
}

type NRCarrierList struct {
	Value []*NRCarrierItem
}

func (ie *NRCarrierList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRCarrierListConstraints)
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

func (ie *NRCarrierList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRCarrierListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NRCarrierItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(NRCarrierItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
