package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceAvailableCapacityConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBPLMNs)),
}

type SliceAvailableCapacity struct {
	Value []*SliceAvailableCapacityItem
}

func (ie *SliceAvailableCapacity) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sliceAvailableCapacityConstraints)
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

func (ie *SliceAvailableCapacity) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sliceAvailableCapacityConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SliceAvailableCapacityItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SliceAvailableCapacityItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
