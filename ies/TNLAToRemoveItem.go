package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLAToRemoveItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tNLAssociationTransportLayerAddress"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type TNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerInformation
	IEExtensions                        []byte
}

func (ie *TNLAToRemoveItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLAToRemoveItemConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Encode(e); err != nil {
		return err
	}
	return nil
}

func (ie *TNLAToRemoveItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLAToRemoveItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Decode(d); err != nil {
		return err
	}
	return nil
}
