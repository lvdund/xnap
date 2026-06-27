package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNSSAIAvailableCapacityListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSliceItems)),
}

type SNSSAIAvailableCapacityList struct {
	Value []*SNSSAIAvailableCapacityItem
}

func (ie *SNSSAIAvailableCapacityList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNSSAIAvailableCapacityListConstraints)
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

func (ie *SNSSAIAvailableCapacityList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNSSAIAvailableCapacityListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SNSSAIAvailableCapacityItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SNSSAIAvailableCapacityItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
