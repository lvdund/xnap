package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsToBeSetupListModifiedSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qfi"},
		{Name: "qosFlowLevelQoSParameters", Optional: true},
		{Name: "offeredGBRQoSFlowInfo", Optional: true},
		{Name: "qosFlowMappingIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsToBeSetupListModifiedSNterminatedItem struct {
	Qfi                       QoSFlowIdentifier
	QosFlowLevelQoSParameters *QoSFlowLevelQoSParameters
	OfferedGBRQoSFlowInfo     *GBRQoSFlowInfo
	QosFlowMappingIndication  *QoSFlowMappingIndication
	IEExtensions              []byte
}

func (ie *QoSFlowsToBeSetupListModifiedSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsToBeSetupListModifiedSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QosFlowLevelQoSParameters != nil, ie.OfferedGBRQoSFlowInfo != nil, ie.QosFlowMappingIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.Qfi.Encode(e); err != nil {
		return err
	}
	if ie.QosFlowLevelQoSParameters != nil {
		if err := ie.QosFlowLevelQoSParameters.Encode(e); err != nil {
			return err
		}
	}
	if ie.OfferedGBRQoSFlowInfo != nil {
		if err := ie.OfferedGBRQoSFlowInfo.Encode(e); err != nil {
			return err
		}
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

func (ie *QoSFlowsToBeSetupListModifiedSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsToBeSetupListModifiedSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Qfi.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.QosFlowLevelQoSParameters = new(QoSFlowLevelQoSParameters)
		if err := ie.QosFlowLevelQoSParameters.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.OfferedGBRQoSFlowInfo = new(GBRQoSFlowInfo)
		if err := ie.OfferedGBRQoSFlowInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
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
