package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CPTransportLayerInformationChEndpointIPAddress = 0
	CPTransportLayerInformationChChoiceExtension   = 1
)

var cPTransportLayerInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "endpointIPAddress"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type CPTransportLayerInformation struct {
	Choice            int
	EndpointIPAddress *TransportLayerAddress
	ChoiceExtension   []byte
}

func (ie *CPTransportLayerInformation) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(cPTransportLayerInformationConstraints)
	switch ie.Choice {
	case 0: // endpointIPAddress
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.EndpointIPAddress.Encode(e); err != nil {
			return err
		}
	case 1: // choice-extension
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *CPTransportLayerInformation) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(cPTransportLayerInformationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // endpointIPAddress
		ie.EndpointIPAddress = new(TransportLayerAddress)
		if err := ie.EndpointIPAddress.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
