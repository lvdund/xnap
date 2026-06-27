package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoSFlowNotifyItemNotificationInformationFulfilled    int64 = 0
	QoSFlowNotifyItemNotificationInformationNotFulfilled int64 = 1
)

var qoSFlowNotifyItemNotificationInformationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type QoSFlowNotifyItemNotificationInformation struct {
	Value int64
}

func (ie *QoSFlowNotifyItemNotificationInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoSFlowNotifyItemNotificationInformationConstraints)
}

func (ie *QoSFlowNotifyItemNotificationInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoSFlowNotifyItemNotificationInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var qoSFlowNotifyItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qosFlowIdentifier"},
		{Name: "notificationInformation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowNotifyItem struct {
	QosFlowIdentifier       QoSFlowIdentifier
	NotificationInformation QoSFlowNotifyItemNotificationInformation
	IEExtensions            []byte
}

func (ie *QoSFlowNotifyItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowNotifyItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.QosFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if err := ie.NotificationInformation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *QoSFlowNotifyItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowNotifyItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if err := ie.NotificationInformation.Decode(d); err != nil {
		return err
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
