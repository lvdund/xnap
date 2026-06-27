package ies

import (
	"github.com/lvdund/asn1go/per"
)

var associatedQoSFlowInfoItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mBS-QoSFlowIdentifier"},
		{Name: "associatedUnicastQoSFlowIdentifier"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AssociatedQoSFlowInfoItem struct {
	MBSQoSFlowIdentifier               QoSFlowIdentifier
	AssociatedUnicastQoSFlowIdentifier QoSFlowIdentifier
	IEExtensions                       []byte
}

func (ie *AssociatedQoSFlowInfoItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(associatedQoSFlowInfoItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.MBSQoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if err := ie.AssociatedUnicastQoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AssociatedQoSFlowInfoItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(associatedQoSFlowInfoItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MBSQoSFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if err := ie.AssociatedUnicastQoSFlowIdentifier.Decode(d); err != nil {
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
