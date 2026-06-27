package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var successfulHOReportInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSuccessfulHOReports)),
}

type SuccessfulHOReportInformation struct {
	Value []*SuccessfulHOReportListItem
}

func (ie *SuccessfulHOReportInformation) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(successfulHOReportInformationConstraints)
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

func (ie *SuccessfulHOReportInformation) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(successfulHOReportInformationConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SuccessfulHOReportListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SuccessfulHOReportListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
