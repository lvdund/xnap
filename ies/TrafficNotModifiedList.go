package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var trafficNotModifiedListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTrafficIndexEntries)),
}

type TrafficNotModifiedList struct {
	Value []*TrafficNotModifiedItem
}

func (ie *TrafficNotModifiedList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(trafficNotModifiedListConstraints)
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

func (ie *TrafficNotModifiedList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(trafficNotModifiedListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TrafficNotModifiedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TrafficNotModifiedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
