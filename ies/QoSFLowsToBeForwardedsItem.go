package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFLowsToBeForwardedsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qosFlowIdentifier"},
		{Name: "dl-dataforwarding"},
		{Name: "ul-dataforwarding"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFLowsToBeForwardedsItem struct {
	QosFlowIdentifier QoSFlowIdentifier
	DlDataforwarding  DLForwarding
	UlDataforwarding  ULForwarding
	IEExtensions      []byte
}

func (ie *QoSFLowsToBeForwardedsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFLowsToBeForwardedsItemConstraints)
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
	if err := ie.DlDataforwarding.Encode(e); err != nil {
		return err
	}
	if err := ie.UlDataforwarding.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *QoSFLowsToBeForwardedsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFLowsToBeForwardedsItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if err := ie.DlDataforwarding.Decode(d); err != nil {
		return err
	}
	if err := ie.UlDataforwarding.Decode(d); err != nil {
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
