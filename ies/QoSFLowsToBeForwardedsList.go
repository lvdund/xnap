package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFLowsToBeForwardedsListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFLowsToBeForwardedsList struct {
	Value []*QoSFLowsToBeForwardedsItem
}

func (ie *QoSFLowsToBeForwardedsList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFLowsToBeForwardedsListConstraints)
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

func (ie *QoSFLowsToBeForwardedsList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFLowsToBeForwardedsListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFLowsToBeForwardedsItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFLowsToBeForwardedsItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
