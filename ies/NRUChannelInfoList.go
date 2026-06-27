package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRUChannelInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofNR_UChannelIDs)),
}

type NRUChannelInfoList struct {
	Value []*NRUChannelInfoItem
}

func (ie *NRUChannelInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRUChannelInfoListConstraints)
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

func (ie *NRUChannelInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRUChannelInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NRUChannelInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(NRUChannelInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
