package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsMappedtoDRBSetupMNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qoSFlowIdentifier"},
		{Name: "qoSFlowLevelQoSParameters"},
		{Name: "qosFlowMappingIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsMappedtoDRBSetupMNterminatedItem struct {
	QoSFlowIdentifier         QoSFlowIdentifier
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters
	QosFlowMappingIndication  *QoSFlowMappingIndication
	IEExtensions              []byte
}

func (ie *QoSFlowsMappedtoDRBSetupMNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsMappedtoDRBSetupMNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QosFlowMappingIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if err := ie.QoSFlowLevelQoSParameters.Encode(e); err != nil {
		return err
	}
	if ie.QosFlowMappingIndication != nil {
		if err := ie.QosFlowMappingIndication.Encode(e); err != nil {
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

func (ie *QoSFlowsMappedtoDRBSetupMNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsMappedtoDRBSetupMNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if err := ie.QoSFlowLevelQoSParameters.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.QosFlowMappingIndication = new(QoSFlowMappingIndication)
		if err := ie.QosFlowMappingIndication.Decode(d); err != nil {
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
