package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var broadcastNIDListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofNIDs)),
}

type BroadcastNIDList struct {
	Value []*BroadcastNIDItem
}

func (ie *BroadcastNIDList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(broadcastNIDListConstraints)
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

func (ie *BroadcastNIDList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(broadcastNIDListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BroadcastNIDItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BroadcastNIDItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
