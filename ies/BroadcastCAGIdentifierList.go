package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var broadcastCAGIdentifierListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCAGs)),
}

type BroadcastCAGIdentifierList struct {
	Value []*BroadcastCAGIdentifierItem
}

func (ie *BroadcastCAGIdentifierList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(broadcastCAGIdentifierListConstraints)
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

func (ie *BroadcastCAGIdentifierList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(broadcastCAGIdentifierListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BroadcastCAGIdentifierItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BroadcastCAGIdentifierItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
