package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFlowNotificationControlIndicationInfoConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFlowNotificationControlIndicationInfo struct {
	Value []*QoSFlowNotifyItem
}

func (ie *QoSFlowNotificationControlIndicationInfo) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFlowNotificationControlIndicationInfoConstraints)
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

func (ie *QoSFlowNotificationControlIndicationInfo) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFlowNotificationControlIndicationInfoConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFlowNotifyItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFlowNotifyItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
