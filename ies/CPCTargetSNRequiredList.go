package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPCTargetSNRequiredListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTargetSNs)),
}

type CPCTargetSNRequiredList struct {
	Value []*CPCTargetSNRequiredListItem
}

func (ie *CPCTargetSNRequiredList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cPCTargetSNRequiredListConstraints)
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

func (ie *CPCTargetSNRequiredList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cPCTargetSNRequiredListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CPCTargetSNRequiredListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CPCTargetSNRequiredListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
