package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var trafficToBeReleaseListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTrafficIndexEntries)),
}

type TrafficToBeReleaseList struct {
	Value []*TrafficToBeReleaseItem
}

func (ie *TrafficToBeReleaseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(trafficToBeReleaseListConstraints)
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

func (ie *TrafficToBeReleaseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(trafficToBeReleaseListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TrafficToBeReleaseItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TrafficToBeReleaseItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
