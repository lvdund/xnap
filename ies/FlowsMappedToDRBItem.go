package ies

import (
	"github.com/lvdund/asn1go/per"
)

var flowsMappedToDRBItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "qoSFlowIdentifier"},
		{Name: "qoSFlowLevelQoSParameters"},
		{Name: "qoSFlowMappingIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type FlowsMappedToDRBItem struct {
	QoSFlowIdentifier         QoSFlowIdentifier
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters
	QoSFlowMappingIndication  *QoSFlowMappingIndication
	IEExtensions              []byte
}

func (ie *FlowsMappedToDRBItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(flowsMappedToDRBItemConstraints)
	if err := seq.EncodePreamble([]bool{ie.QoSFlowMappingIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if err := ie.QoSFlowLevelQoSParameters.Encode(e); err != nil {
		return err
	}
	if ie.QoSFlowMappingIndication != nil {
		if err := ie.QoSFlowMappingIndication.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *FlowsMappedToDRBItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(flowsMappedToDRBItemConstraints)
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
		ie.QoSFlowMappingIndication = new(QoSFlowMappingIndication)
		if err := ie.QoSFlowMappingIndication.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
