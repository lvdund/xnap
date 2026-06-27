package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFLowsAcceptedToBeForwardedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qosFlowIdentifier"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFLowsAcceptedToBeForwardedItem struct {
	QosFlowIdentifier QoSFlowIdentifier
	IEExtensions      []byte
}

func (ie *QoSFLowsAcceptedToBeForwardedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFLowsAcceptedToBeForwardedItemConstraints)
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
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *QoSFLowsAcceptedToBeForwardedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFLowsAcceptedToBeForwardedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosFlowIdentifier.Decode(d); err != nil {
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
