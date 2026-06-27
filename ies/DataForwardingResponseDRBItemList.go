package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dataForwardingResponseDRBItemListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DataForwardingResponseDRBItemList struct {
	Value []*DataForwardingResponseDRBItem
}

func (ie *DataForwardingResponseDRBItemList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dataForwardingResponseDRBItemListConstraints)
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

func (ie *DataForwardingResponseDRBItemList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dataForwardingResponseDRBItemListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DataForwardingResponseDRBItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DataForwardingResponseDRBItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
