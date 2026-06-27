package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dataforwardingandOffloadingInfofromSourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qosFlowsToBeForwarded"},
		{Name: "sourceDRBtoQoSFlowMapping", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DataforwardingandOffloadingInfofromSource struct {
	QosFlowsToBeForwarded     QoSFLowsToBeForwardedsList
	SourceDRBtoQoSFlowMapping *DRBToQoSFlowMappingList
	IEExtensions              []byte
}

func (ie *DataforwardingandOffloadingInfofromSource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataforwardingandOffloadingInfofromSourceConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SourceDRBtoQoSFlowMapping != nil, false}); err != nil {
		return err
	}
	if err := ie.QosFlowsToBeForwarded.Encode(e); err != nil {
		return err
	}
	if ie.SourceDRBtoQoSFlowMapping != nil {
		if err := ie.SourceDRBtoQoSFlowMapping.Encode(e); err != nil {
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

func (ie *DataforwardingandOffloadingInfofromSource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataforwardingandOffloadingInfofromSourceConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosFlowsToBeForwarded.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SourceDRBtoQoSFlowMapping = new(DRBToQoSFlowMappingList)
		if err := ie.SourceDRBtoQoSFlowMapping.Decode(d); err != nil {
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
