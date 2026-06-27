package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var broadcastPNINPNIDInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBPLMNs)),
}

type BroadcastPNINPNIDInformation struct {
	Value []*BroadcastPNINPNIDInformationItem
}

func (ie *BroadcastPNINPNIDInformation) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(broadcastPNINPNIDInformationConstraints)
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

func (ie *BroadcastPNINPNIDInformation) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(broadcastPNINPNIDInformationConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BroadcastPNINPNIDInformationItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BroadcastPNINPNIDInformationItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
