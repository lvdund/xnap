package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var coverageModificationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(0)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
}

type CoverageModificationList struct {
	Value []*CoverageModificationListItem
}

func (ie *CoverageModificationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(coverageModificationListConstraints)
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

func (ie *CoverageModificationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(coverageModificationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CoverageModificationListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CoverageModificationListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
