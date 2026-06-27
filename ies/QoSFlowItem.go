package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qfi"},
		{Name: "qosFlowMappingIndication", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowItem struct {
	Qfi                      QoSFlowIdentifier
	QosFlowMappingIndication *QoSFlowMappingIndication
	IEExtensions             []byte
}

func (ie *QoSFlowItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QosFlowMappingIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.Qfi.Encode(e); err != nil {
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

func (ie *QoSFlowItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowItemConstraints)
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
