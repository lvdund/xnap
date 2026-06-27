package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBToReportListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
}

type SSBToReportList struct {
	Value []*SSBToReportListItem
}

func (ie *SSBToReportList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBToReportListConstraints)
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

func (ie *SSBToReportList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBToReportListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SSBToReportListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SSBToReportListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
