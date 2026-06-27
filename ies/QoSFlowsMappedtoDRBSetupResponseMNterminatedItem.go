package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsMappedtoDRBSetupResponseMNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qoSFlowIdentifier"},
		{Name: "currentQoSParaSetIndex"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsMappedtoDRBSetupResponseMNterminatedItem struct {
	QoSFlowIdentifier      QoSFlowIdentifier
	CurrentQoSParaSetIndex QoSParaSetIndex
	IEExtensions           []byte
}

func (ie *QoSFlowsMappedtoDRBSetupResponseMNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsMappedtoDRBSetupResponseMNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if err := ie.CurrentQoSParaSetIndex.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *QoSFlowsMappedtoDRBSetupResponseMNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsMappedtoDRBSetupResponseMNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if err := ie.CurrentQoSParaSetIndex.Decode(d); err != nil {
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
