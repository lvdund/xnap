package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourcesNotifyItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "qosFlowsNotificationContrIndInfo"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourcesNotifyItem struct {
	PduSessionId                     PDUSessionID
	QosFlowsNotificationContrIndInfo QoSFlowNotificationControlIndicationInfo
	IEExtensions                     []byte
}

func (ie *PDUSessionResourcesNotifyItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourcesNotifyItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if err := ie.QosFlowsNotificationContrIndInfo.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionResourcesNotifyItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourcesNotifyItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if err := ie.QosFlowsNotificationContrIndInfo.Decode(d); err != nil {
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
