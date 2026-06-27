package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var trafficRequiredModifiedListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTrafficIndexEntries)),
}

type TrafficRequiredModifiedList struct {
	Value []*TrafficRequiredModifiedItem
}

func (ie *TrafficRequiredModifiedList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(trafficRequiredModifiedListConstraints)
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

func (ie *TrafficRequiredModifiedList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(trafficRequiredModifiedListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TrafficRequiredModifiedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TrafficRequiredModifiedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
