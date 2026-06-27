package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPCTargetSNConfirmListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTargetSNs)),
}

type CPCTargetSNConfirmList struct {
	Value []*CPCTargetSNConfirmListItem
}

func (ie *CPCTargetSNConfirmList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cPCTargetSNConfirmListConstraints)
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

func (ie *CPCTargetSNConfirmList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cPCTargetSNConfirmListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CPCTargetSNConfirmListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CPCTargetSNConfirmListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
