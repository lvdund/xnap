package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLAFailedToSetupItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tNLAssociationTransportLayerAddress"},
		{Name: "cause"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type TNLAFailedToSetupItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerInformation
	Cause                               Cause
	IEExtensions                        []byte
}

func (ie *TNLAFailedToSetupItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLAFailedToSetupItemConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Encode(e); err != nil {
		return err
	}
	if err := ie.Cause.Encode(e); err != nil {
		return err
	}
	return nil
}

func (ie *TNLAFailedToSetupItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLAFailedToSetupItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TNLAssociationTransportLayerAddress.Decode(d); err != nil {
		return err
	}
	if err := ie.Cause.Decode(d); err != nil {
		return err
	}
	return nil
}
