package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cHOTargetSNNodeListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTargetSNs)),
}

type CHOTargetSNNodeList struct {
	Value []*CHOTargetSNNodeItem
}

func (ie *CHOTargetSNNodeList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cHOTargetSNNodeListConstraints)
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

func (ie *CHOTargetSNNodeList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cHOTargetSNNodeListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CHOTargetSNNodeItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CHOTargetSNNodeItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
