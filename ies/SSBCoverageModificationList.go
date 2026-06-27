package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBCoverageModificationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(0)),
	Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
}

type SSBCoverageModificationList struct {
	Value []*SSBCoverageModificationListItem
}

func (ie *SSBCoverageModificationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBCoverageModificationListConstraints)
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

func (ie *SSBCoverageModificationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBCoverageModificationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SSBCoverageModificationListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SSBCoverageModificationListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
