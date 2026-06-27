package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLAToAddItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tNLAssociationTransportLayerAddress"},
		{Name: "tNLAssociationUsage"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type TNLAToAddItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerInformation
	TNLAssociationUsage                 TNLAssociationUsage
	IEExtensions                        []byte
}

func (ie *TNLAToAddItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLAToAddItemConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Encode(e); err != nil {
		return err
	}
	if err := ie.TNLAssociationUsage.Encode(e); err != nil {
		return err
	}
	return nil
}

func (ie *TNLAToAddItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLAToAddItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Decode(d); err != nil {
		return err
	}
	if err := ie.TNLAssociationUsage.Decode(d); err != nil {
		return err
	}
	return nil
}
