package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourcesActivityNotifyItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "pduSessionLevelUPactivityreport", Optional: true},
		{Name: "qosFlowsActivityNotifyList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourcesActivityNotifyItem struct {
	PduSessionId                    PDUSessionID
	PduSessionLevelUPactivityreport *UserPlaneTrafficActivityReport
	QosFlowsActivityNotifyList      *QoSFlowsActivityNotifyList
	IEExtensions                    []byte
}

func (ie *PDUSessionResourcesActivityNotifyItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourcesActivityNotifyItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionLevelUPactivityreport != nil, ie.QosFlowsActivityNotifyList != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if ie.PduSessionLevelUPactivityreport != nil {
		if err := ie.PduSessionLevelUPactivityreport.Encode(e); err != nil {
			return err
		}
	}
	if ie.QosFlowsActivityNotifyList != nil {
		if err := ie.QosFlowsActivityNotifyList.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionResourcesActivityNotifyItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourcesActivityNotifyItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PduSessionLevelUPactivityreport = new(UserPlaneTrafficActivityReport)
		if err := ie.PduSessionLevelUPactivityreport.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.QosFlowsActivityNotifyList = new(QoSFlowsActivityNotifyList)
		if err := ie.QosFlowsActivityNotifyList.Decode(d); err != nil {
			return err
		}
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
