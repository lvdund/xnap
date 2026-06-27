package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFlowsUsageReportListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFlowsUsageReportList struct {
	Value []*QoSFlowsUsageReportItem
}

func (ie *QoSFlowsUsageReportList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFlowsUsageReportListConstraints)
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

func (ie *QoSFlowsUsageReportList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFlowsUsageReportListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFlowsUsageReportItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFlowsUsageReportItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
