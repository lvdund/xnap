package ies

import (
	"github.com/lvdund/asn1go/per"
)

var secondaryRATUsageInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pDUSessionUsageReport", Optional: true},
		{Name: "qosFlowsUsageReportList", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SecondaryRATUsageInformation struct {
	PDUSessionUsageReport   *PDUSessionUsageReport
	QosFlowsUsageReportList *QoSFlowsUsageReportList
	IEExtensions            []byte
}

func (ie *SecondaryRATUsageInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(secondaryRATUsageInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PDUSessionUsageReport != nil, ie.QosFlowsUsageReportList != nil, false}); err != nil {
		return err
	}
	if ie.PDUSessionUsageReport != nil {
		if err := ie.PDUSessionUsageReport.Encode(e); err != nil {
			return err
		}
	}
	if ie.QosFlowsUsageReportList != nil {
		if err := ie.QosFlowsUsageReportList.Encode(e); err != nil {
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

func (ie *SecondaryRATUsageInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(secondaryRATUsageInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PDUSessionUsageReport = new(PDUSessionUsageReport)
		if err := ie.PDUSessionUsageReport.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.QosFlowsUsageReportList = new(QoSFlowsUsageReportList)
		if err := ie.QosFlowsUsageReportList.Decode(d); err != nil {
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
