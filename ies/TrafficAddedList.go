package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var trafficAddedListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTrafficIndexEntries)),
}

type TrafficAddedList struct {
	Value []*TrafficAddedItem
}

func (ie *TrafficAddedList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(trafficAddedListConstraints)
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

func (ie *TrafficAddedList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(trafficAddedListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TrafficAddedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TrafficAddedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
