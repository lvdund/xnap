package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLAToUpdateItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tNLAssociationTransportLayerAddress"},
		{Name: "tNLAssociationUsage", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type TNLAToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerInformation
	TNLAssociationUsage                 *TNLAssociationUsage
	IEExtensions                        []byte
}

func (ie *TNLAToUpdateItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLAToUpdateItemConstraints)
	if err := seq.EncodePreamble([]bool{ie.TNLAssociationUsage != nil, false}); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Encode(e); err != nil {
		return err
	}
	if ie.TNLAssociationUsage != nil {
		if err := ie.TNLAssociationUsage.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TNLAToUpdateItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLAToUpdateItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.TNLAssociationUsage = new(TNLAssociationUsage)
		if err := ie.TNLAssociationUsage.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
