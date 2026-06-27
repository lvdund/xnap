package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsMappedtoDRBSetupResponseSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qoSFlowIdentifier"},
		{Name: "mCGRequestedGBRQoSFlowInfo", Optional: true},
		{Name: "qosFlowMappingIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsMappedtoDRBSetupResponseSNterminatedItem struct {
	QoSFlowIdentifier          QoSFlowIdentifier
	MCGRequestedGBRQoSFlowInfo *GBRQoSFlowInfo
	QosFlowMappingIndication   *QoSFlowMappingIndication
	IEExtensions               []byte
}

func (ie *QoSFlowsMappedtoDRBSetupResponseSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsMappedtoDRBSetupResponseSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MCGRequestedGBRQoSFlowInfo != nil, ie.QosFlowMappingIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if ie.MCGRequestedGBRQoSFlowInfo != nil {
		if err := ie.MCGRequestedGBRQoSFlowInfo.Encode(e); err != nil {
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

func (ie *QoSFlowsMappedtoDRBSetupResponseSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsMappedtoDRBSetupResponseSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MCGRequestedGBRQoSFlowInfo = new(GBRQoSFlowInfo)
		if err := ie.MCGRequestedGBRQoSFlowInfo.Decode(d); err != nil {
			return err
		}
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
