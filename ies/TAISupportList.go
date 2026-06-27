package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tAISupportListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofsupportedTACs)),
}

type TAISupportList struct {
	Value []*TAISupportItem
}

func (ie *TAISupportList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tAISupportListConstraints)
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

func (ie *TAISupportList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tAISupportListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TAISupportItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TAISupportItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
