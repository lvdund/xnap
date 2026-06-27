package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoSFlowsUsageReportItemRATTypeNr              int64 = 0
	QoSFlowsUsageReportItemRATTypeEutra           int64 = 1
	QoSFlowsUsageReportItemRATTypeNrUnlicensed    int64 = 2
	QoSFlowsUsageReportItemRATTypeEUtraUnlicensed int64 = 3
)

var qoSFlowsUsageReportItemRATTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2, 3},
}

type QoSFlowsUsageReportItemRATType struct {
	Value int64
}

func (ie *QoSFlowsUsageReportItemRATType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoSFlowsUsageReportItemRATTypeConstraints)
}

func (ie *QoSFlowsUsageReportItemRATType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoSFlowsUsageReportItemRATTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var qoSFlowsUsageReportItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qosFlowIdentifier"},
		{Name: "rATType"},
		{Name: "qoSFlowsTimedReportList"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       QoSFlowIdentifier
	RATType                 QoSFlowsUsageReportItemRATType
	QoSFlowsTimedReportList VolumeTimedReportList
	IEExtensions            []byte
}

func (ie *QoSFlowsUsageReportItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsUsageReportItemConstraints)
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
	if err := ie.RATType.Encode(e); err != nil {
		return err
	}
	if err := ie.QoSFlowsTimedReportList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *QoSFlowsUsageReportItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsUsageReportItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if err := ie.RATType.Decode(d); err != nil {
		return err
	}
	if err := ie.QoSFlowsTimedReportList.Decode(d); err != nil {
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
