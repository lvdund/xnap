package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLASetupItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tNLAssociationTransportLayerAddress"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TNLASetupItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerInformation
	IEExtensions                        []byte
}

func (ie *TNLASetupItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLASetupItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TNLASetupItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLASetupItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Decode(d); err != nil {
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
