package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bAPControlPDURLCCHListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBAPControlPDURLCCHs)),
}

type BAPControlPDURLCCHList struct {
	Value []*BAPControlPDURLCCHItem
}

func (ie *BAPControlPDURLCCHList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bAPControlPDURLCCHListConstraints)
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

func (ie *BAPControlPDURLCCHList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bAPControlPDURLCCHListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*BAPControlPDURLCCHItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(BAPControlPDURLCCHItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
