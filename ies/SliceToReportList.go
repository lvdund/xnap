package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceToReportListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBPLMNs)),
}

type SliceToReportList struct {
	Value []*SliceToReportListItem
}

func (ie *SliceToReportList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sliceToReportListConstraints)
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

func (ie *SliceToReportList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sliceToReportListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SliceToReportListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SliceToReportListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
