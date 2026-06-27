package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFLowsAcceptedToBeForwardedListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFLowsAcceptedToBeForwardedList struct {
	Value []*QoSFLowsAcceptedToBeForwardedItem
}

func (ie *QoSFLowsAcceptedToBeForwardedList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFLowsAcceptedToBeForwardedListConstraints)
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

func (ie *QoSFLowsAcceptedToBeForwardedList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFLowsAcceptedToBeForwardedListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFLowsAcceptedToBeForwardedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFLowsAcceptedToBeForwardedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
